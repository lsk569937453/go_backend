package dao

import (
	"fmt"
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
		fmt.Printf("query failed, err:%v\n", err)
	}

	// var users []vojo.TasksDao
	// err = nstmt.Select(&users, map[string]interface{}{"user_id": "-1"})
	// if err != nil {
	// 	fmt.Printf("query failed, err:%v\n", err)

	// }
}
func HistoryGetById(req *vojo.GetTaskHistoryByTaskIdReq) []vojo.TasksHistory {
	sqlStr := "SELECT id,task_id, exec_time , exec_result,exec_code,_timestamp FROM task_exec_history where task_id=? order by id desc limit 100"
	var taskHistory []vojo.TasksHistory
	err := CronDb.Select(&taskHistory, sqlStr, req.TaskId)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)

	}
	return taskHistory
}
