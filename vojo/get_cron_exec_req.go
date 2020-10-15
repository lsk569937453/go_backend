package vojo

type GetCronExecReq struct {
	CronList []string `form:"cronList" json:"cronList" binding:"required"`
}
