package dao

import (
	"encoding/base64"
	"go_backend/log"
	"go_backend/vojo"
	"reflect"
)

func HistoryInsert(history vojo.TasksHistory) {
	params := make(map[string]interface{})

	elem := reflect.ValueOf(&history).Elem()
	relType := elem.Type()
	for i := 0; i < relType.NumField(); i++ {
		params[relType.Field(i).Name] = elem.Field(i).Interface()
	}

	_, err := CronDb.NamedExec(`insert into task_exec_history ( task_id, exec_time , exec_result,exec_code)values(:Task_id,:Exec_time,:Exec_result,:Exec_code)`, params)
	if err != nil {
		log.Error("HistoryInsert failed, err:%v", err.Error())
	}


}
func HistoryGetById(req *vojo.GetTaskHistoryByTaskIdReq) []vojo.TasksHistory {
	sqlStr := "SELECT id,task_id, exec_time , exec_result,exec_code,_timestamp FROM task_exec_history where task_id=? order by id desc limit 100"
	var taskHistory []vojo.TasksHistory
	err := CronDb.Select(&taskHistory, sqlStr, req.TaskId)
	if err != nil {
		log.Errorf("query failed, err:%v", err.Error())

	}
	if taskHistory==nil{
		taskHistory=make([]vojo.TasksHistory,0)
	}else {
		for i,item:=range  taskHistory{
			decodeBytes, err := base64.StdEncoding.DecodeString(item.Exec_result)
			if err!=nil{
				log.Errorf("base64 decode %s,historyId:%s",err.Error(),item.Id)
			}else {
				taskHistory[i].Exec_result=string(decodeBytes)
			}
		}

	}
	return taskHistory
}
