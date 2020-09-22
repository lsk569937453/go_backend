package dao

import (
	"encoding/base64"
	"go_backend/log"
	"go_backend/vojo"
	"reflect"
)

func HistoryInsert(history *vojo.TasksHistory) error {
	err := history.TaskHistoryInsertValidator()
	if err != nil {
		return err
	}

	params := make(map[string]interface{})

	elem := reflect.ValueOf(history).Elem()
	relType := elem.Type()
	for i := 0; i < relType.NumField(); i++ {
		params[relType.Field(i).Name] = elem.Field(i).Interface()
	}

	_, err = CronDb.NamedExec(`insert into task_exec_history ( task_id, exec_time , exec_result,exec_code)values(:Task_id,:Exec_time,:Exec_result,:Exec_code)`, params)
	if err != nil {
		log.Error("HistoryInsert failed, err:%v", err.Error())
		return err
	}
	return nil

}

/**
 *
 * @Description  getAllHistoryById get  all the history id
 * @Date 10:27 上午 2020/9/3
 **/
func HistoryGetById(req *vojo.GetTaskHistoryByTaskIdReq) []vojo.TasksHistory {
	sqlStr := "SELECT id,task_id, exec_time , exec_result,exec_code,_timestamp FROM task_exec_history where task_id=? order by id desc limit 100"
	var taskHistory []vojo.TasksHistory
	err := CronDb.Select(&taskHistory, sqlStr, req.TaskId)
	if err != nil {
		log.Errorf("query failed, err:%v", err.Error())

	}
	if taskHistory == nil {
		taskHistory = make([]vojo.TasksHistory, 0)
	} else {
		for i, item := range taskHistory {
			decodeBytes, err := base64.StdEncoding.DecodeString(item.Exec_result)
			if err != nil {
				log.Errorf("base64 decode %s,historyId:%s", err.Error(), item.Id)
			} else {
				taskHistory[i].Exec_result = string(decodeBytes)
			}
		}

	}
	return taskHistory
}

/**
 *
 * @Description  getHistoryDataByPage
 * @Date 1:53 下午 2020/9/3
 **/
func HistotyGetByPage(req *vojo.GetHistoryByPage) []vojo.TasksHistory {
	var taskHistory []vojo.TasksHistory
	var err error
	defaultSql := "SELECT id,task_id, exec_time , exec_result,exec_code,_timestamp FROM task_exec_history where task_id=? order by id desc limit 20"
	if req.Page != nil {
		defaultSql = "SELECT id,task_id, exec_time , exec_result,exec_code,_timestamp FROM task_exec_history where task_id=? and id <? order by id desc limit ?"
		err = CronDb.Select(&taskHistory, defaultSql, req.TaskId, req.Page.Id, req.Page.PageSize)

	} else {
		err = CronDb.Select(&taskHistory, defaultSql, req.TaskId)
	}
	if err != nil {
		log.Error("history get by page error:%s", err.Error())
	}
	if taskHistory == nil {
		taskHistory = make([]vojo.TasksHistory, 0)
	} else {
		for i, item := range taskHistory {
			decodeBytes, err := base64.StdEncoding.DecodeString(item.Exec_result)
			if err != nil {
				log.Errorf("base64 decode %s,historyId:%s", err.Error(), item.Id)
			} else {
				taskHistory[i].Exec_result = string(decodeBytes)
			}
		}

	}
	return taskHistory

}

/**
 *
 * @Description  get the row count
 * @Date 2:31 下午 2020/8/31
 **/
func HitoryCount() int64 {
	sqlStr := "select count(*) from task_exec_history"
	var count int64
	rows := CronDb.QueryRow(sqlStr)

	err := rows.Scan(&count)
	if err != nil {
		log.Error("del task error:%v", err.Error())
		return -1
	}
	return count

}

/**
 *
 * @Description  delete the last n row data
 * @Date 2:22 下午 2020/8/31
 **/
func HitoryDeleteLast(rowcount int) {
	sqlStr := "DELETE FROM task_exec_history ORDER BY  id asc limit ?"
	_, err := CronDb.Exec(sqlStr, rowcount)
	if err != nil {
		log.Error("del task error:%v", err.Error())
	}
}
