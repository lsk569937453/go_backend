package dao

import (
	"fmt"
	"go_backend/vojo"

)

func Insert() {
	var params map[string]interface{}
	params = map[string]interface{}{"user_id": "-1"}
	_, err := CronDb.NamedExec(`insert into tasks ( task_cron, task_name , user_id,url)values(:task_cron,:task_name,:user_id,:url)`, params)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
	}

	// var users []vojo.TasksDao
	// err = nstmt.Select(&users, map[string]interface{}{"user_id": "-1"})
	// if err != nil {
	// 	fmt.Printf("query failed, err:%v\n", err)

	// }
}
func GetTaskById() []vojo.TasksDao {
	sqlStr := "SELECT id,task_name, task_cron, url,user_id FROM tasks"
	var users []vojo.TasksDao
	err := CronDb.Select(&users, sqlStr)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)

	}
	return users
}

func GetTaskByUserId() []vojo.TasksDao {

	nstmt, err := CronDb.PrepareNamed(`SELECT id,task_name,task_cron,url,_timestamp FROM tasks where user_id=:user_id`)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
	}

	var users []vojo.TasksDao
	err = nstmt.Select(&users, map[string]interface{}{"user_id": "-1"})
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)

	}
	return users
}
