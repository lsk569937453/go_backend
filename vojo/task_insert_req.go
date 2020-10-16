package vojo

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/robfig/cron/v3"
	"github.com/valyala/fasthttp"
	"go_backend/log"
	"time"
)

type TaskInsertReq struct {
	Name           string `form:"name" json:"name" validate:"required"`
	CronExpression string `form:"cron_expression"  json:"cron_expression" validate:"cronExpression"`
	Url            string `form:"url" json:"url" validate:"url"`
}

func ValidateCronExpression(f1 validator.FieldLevel) bool { //验证字段的方法的定义
	if f1.Field().String() == "" {
		return false
	}
	parser := cron.NewParser(
		cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor,
	)

	s, err := parser.Parse(f1.Field().String())
	if err != nil {
		return false
	}
	theTime := s.Next(time.Now()).Format("2006-01-02 15:04:05")
	fmt.Println(theTime)
	return true

}
func ValidateUrl(f1 validator.FieldLevel) bool { //验证字段的方法的定义
	url := f1.Field().String()
	_, resp, err := fasthttp.Get(nil, url)
	var responseBody string
	if err != nil {
		responseBody = err.Error()

		log.Error("doReq error,%s", responseBody)
		return false
	} else {
		responseBody = string(resp)
	}
	fmt.Println(responseBody)
	return true
}

func (u *TaskInsertReq) UserValidator() error { //自定义的验证函数，
	validata := validator.New()
	err := validata.RegisterValidation("cronExpression", ValidateCronExpression) //注册验证字段和字段验证的功能
	if err != nil {
		return err
	}
	err = validata.RegisterValidation("url", ValidateUrl)
	if err != nil {
		return err
	}
	err = validata.Struct(u)
	return err

}
