package vojo

type CronResult struct {
	Result         []string `form:"cronList" json:"result"`
	CronExpression string   `form:"cronList" json:"cronExpression"`
	Error          string   `form:"cronList" json:"error"`
}
