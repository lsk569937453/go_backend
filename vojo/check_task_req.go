package vojo

type CheckTaskReq struct {
	Name           string `form:"name" binding:"required"`
	CronExpression string `form:"cron_expression" binding:"required"`
}
