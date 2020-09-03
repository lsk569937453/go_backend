package vojo

type GetHistoryByPage struct {
	TaskId int         `form:"task_id" json:"task_id" `
	Page   *PageHelper `form:"page" json:"page" `
}
