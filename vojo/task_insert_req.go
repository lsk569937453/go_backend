package vojo
type TaskInsertReq struct {
	Name           string `form:"name" json:"name" `
	CronExpression string `form:"cron_expression"  json:"cron_expression"`
	Url string `form:"url" json:"url"`
}
