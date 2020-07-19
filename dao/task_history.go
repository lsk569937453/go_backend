package dao

import (
	"fmt"
	"go_backend/vojo"

)

func HistoryInsert() {
	var params map[string]interface{}
	params = map[string]interface{}{"user_id": "-1"}
	_, err := CronDb.NamedExec(`insert into task_exec_history ( task_id, exec_time , exec_result,exec_code)values(:task_id,:exec_time,:exec_result,:exec_code)`, params)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
	}

	// var users []vojo.TasksDao
	// err = nstmt.Select(&users, map[string]interface{}{"user_id": "-1"})
	// if err != nil {
	// 	fmt.Printf("query failed, err:%v\n", err)

	// }
}
func HistoryGetById() []vojo.TasksHistory {
	sqlStr := "SELECT id,task_id, exec_time , exec_result,exec_code,_timestamp FROM task_exec_history"
	var taskHistory []vojo.TasksHistory
	err := CronDb.Select(&taskHistory, sqlStr)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)

	}
	return taskHistory
}
