package vojo

type TaskUpdateReq struct {
	Id             int    `form:"id" json:"id" `
	CronExpression string `form:"cron_expression"  json:"cron_expression"`
	Url            string `form:"url" json:"url"`
}
