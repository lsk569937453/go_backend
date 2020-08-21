package dao

import (
	"fmt"
	"go_backend/log"
	"go_backend/vojo"

	"reflect"
)

func AddTask(req vojo.TaskInsertReq) vojo.BaseRes{
	 params:=make(map[string]interface{})

	elem := reflect.ValueOf(&req).Elem()
	relType := elem.Type()
	for i := 0; i < relType.NumField(); i++ {
		params[relType.Field(i).Name] = elem.Field(i).Interface()
	}
	params["user_id"]=-1
	_, err := CronDb.NamedExec(`insert into tasks ( task_cron, task_name , user_id,url)values(:CronExpression,:Name,:user_id,:Url)`, params)
	if err != nil {
		log.Error("query failed, err:%v\n", err)
	}
	var res vojo.BaseRes
	return  res

	// var users []vojo.TasksDao
	// err = nstmt.Select(&users, map[string]interface{}{"user_id": "-1"})
	// if err != nil {
	// 	fmt.Printf("query failed, err:%v\n", err)

	// }
}
func GetTaskByUserId(req*vojo.GetTaskByUserIdReq) []vojo.TasksDao {
	sqlStr := "SELECT id,task_name, task_cron, url,user_id FROM tasks where user_id=?"
	var users []vojo.TasksDao
	err := CronDb.Select(&users, sqlStr,req.UserId)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
	}
	return users
}
func GetTaskById(req*vojo.GetTaskByIdReq) []vojo.TasksDao {
	sqlStr := "SELECT id,task_name, task_cron, url,user_id FROM tasks where id=?"
	var users []vojo.TasksDao
	err := CronDb.Select(&users, sqlStr,req.Id)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
	}
	return users
}
func GetAllTask() []vojo.TasksDao {
	sqlStr := "SELECT id,task_name, task_cron, url,user_id FROM tasks"
	var users []vojo.TasksDao
	err := CronDb.Select(&users, sqlStr)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
	}
	return users
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
