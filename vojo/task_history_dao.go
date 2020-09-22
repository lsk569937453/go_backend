package vojo

import (
	"github.com/go-playground/validator/v10"
)

type TasksHistory struct {
	Task_id     int    `db:"task_id" json:"task_id" validate:"TaskId"`
	Exec_time   string `db:"exec_time" json:"exec_time" validate:"ExecTime"`
	Exec_result string `db:"exec_result" json:"exec_result" validate:"ExecResult"`
	Exec_code   int    `db:"exec_code" json:"exec_code" `
	Id          int    `db:"id" json:"id"`
	Timestamp   string `db:"_timestamp" json:"_timestamp"`
}

func ValidateTaskId(f1 validator.FieldLevel) bool { //验证字段的方法的定义
	if f1.Field().Int() == 0 {
		return false
	}
	return true
}
func ValidateExecTime(f1 validator.FieldLevel) bool { //验证字段的方法的定义
	if f1.Field().String() == "" {
		return false
	}
	return true
}
func ValidateExecResult(f1 validator.FieldLevel) bool { //验证字段的方法的定义
	if f1.Field().String() == "" {
		return false
	}
	return true
}

func (u *TasksHistory) TaskHistoryInsertValidator() error { //自定义的验证函数，
	validata := validator.New()
	validata.RegisterValidation("TaskId", ValidateTaskId) //注册验证字段和字段验证的功能
	validata.RegisterValidation("ExecTime", ValidateExecTime)
	validata.RegisterValidation("ExecResult", ValidateExecResult)
	err := validata.Struct(u)
	return err

}
