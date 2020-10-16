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
	return f1.Field().Int() != 0
}
func ValidateExecTime(f1 validator.FieldLevel) bool { //验证字段的方法的定义
	return f1.Field().String() != ""
}
func ValidateExecResult(f1 validator.FieldLevel) bool { //验证字段的方法的定义
	return f1.Field().String() != ""

}

func (u *TasksHistory) TaskHistoryInsertValidator() error { //自定义的验证函数，
	validata := validator.New()
	err := validata.RegisterValidation("TaskId", ValidateTaskId) //注册验证字段和字段验证的功能
	if err != nil {
		return err
	}
	err = validata.RegisterValidation("ExecTime", ValidateExecTime)
	if err != nil {
		return err
	}
	err = validata.RegisterValidation("ExecResult", ValidateExecResult)
	if err != nil {
		return err
	}
	err = validata.Struct(u)
	return err

}
