package dao

import (
	"go_backend/log"
	"go_backend/vojo"

	"reflect"
)

/**
 *
 * @Description
 * @Date 2:35 下午 2020/8/24
 **/
func AddTask(req vojo.TaskInsertReq) int64 {
	params := make(map[string]interface{})

	elem := reflect.ValueOf(&req).Elem()
	relType := elem.Type()
	for i := 0; i < relType.NumField(); i++ {
		params[relType.Field(i).Name] = elem.Field(i).Interface()
	}
	params["user_id"] = -1
	result, err := CronDb.NamedExec(`insert into tasks ( task_cron, task_name , user_id,url)values(:CronExpression,:Name,:user_id,:Url)`, params)
	if err != nil {
		log.Error("query failed, err:%v", err.Error())
		return -1

	} else {
		insertId, err := result.LastInsertId()
		if err != nil {
			log.Error("LastInsertId failed, err:%v", err.Error())
			return -1
		} else {
			return insertId
			//task.AddTask(req.CronExpression, req.Url, int(insertId))
		}
	}

	// var users []vojo.TasksDao
	// err = nstmt.Select(&users, map[string]interface{}{"user_id": "-1"})
	// if err != nil {
	// 	fmt.Printf("query failed, err:%v\n", err)

	// }
}

/**
 *
 * @Description //TODO
 * @Date 2:29 下午 2020/8/24
 * @Param
 * @return
 **/
func GetTaskByUserId(req *vojo.GetTaskByUserIdReq) []vojo.TasksDao {
	sqlStr := "SELECT id,task_name, task_cron, url,user_id,_timestamp FROM tasks where user_id=?"
	var users []vojo.TasksDao
	err := CronDb.Select(&users, sqlStr, req.UserId)
	if err != nil {
		log.Errorf("query failed, err:%v", err.Error())
	}
	return users
}
func GetTaskById(req *vojo.GetTaskByIdReq) []vojo.TasksDao {
	sqlStr := "SELECT id,task_name, task_cron, url,user_id ,_timestamp FROM tasks where id=?"
	var users []vojo.TasksDao
	err := CronDb.Select(&users, sqlStr, req.Id)
	if err != nil {
		log.Errorf("query failed, err:%v", err.Error())
	}
	return users
}
func GetAllTask() []vojo.TasksDao {
	sqlStr := "SELECT id,task_name, task_cron, url,user_id FROM tasks"
	var users []vojo.TasksDao
	err := CronDb.Select(&users, sqlStr)
	if err != nil {
		log.Errorf("query failed, err:%v", err.Error())
	}
	return users
}
func UpdateTask(req vojo.TaskUpdateReq) int {
	var err error
	if req.Url != "" && req.CronExpression != "" {
		sqlStr := "update tasks set  task_cron=? , url=? where id=?"
		_, err = CronDb.Exec(sqlStr, req.CronExpression, req.Url, req.Id)
	} else if req.Url != "" && req.CronExpression == "" {
		sqlStr := "update tasks set   url=? where id=?"
		_, err = CronDb.Exec(sqlStr, req.Url, req.Id)
	} else if req.Url == "" && req.CronExpression != "" {
		sqlStr := "update tasks set  task_cron=? where id=?"
		_, err = CronDb.Exec(sqlStr, req.CronExpression, req.Id)
	} else {
		log.Error("update task error")
	}
	if err != nil {
		log.Errorf("update task error:%v", err.Error())
		return -1
	}

	return 0
}
func DelTask(req vojo.TaskDelByIdReq) int {
	sqlStr := "delete from tasks  where id=?"
	_, err := CronDb.Exec(sqlStr, req.Id)
	if err != nil {
		log.Error("del task error:%v", err.Error())
		return -1
	}

	return 0
}

//func GetTaskByUserId() []vojo.TasksDao {
//
//	nstmt, err := CronDb.PrepareNamed(`SELECT id,task_name,task_cron,url,_timestamp FROM tasks where user_id=:user_id`)
//	if err != nil {
//		fmt.Printf("query failed, err:%v\n", err)
//	}
//
//	var users []vojo.TasksDao
//	err = nstmt.Select(&users, map[string]interface{}{"user_id": "-1"})
//	if err != nil {
//		fmt.Printf("query failed, err:%v\n", err)
//
//	}
//	return users
//}
