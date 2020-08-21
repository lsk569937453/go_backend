package vojo

type TasksDao struct {
	Task_name string `db:"task_name" json:"task_name"`
	Task_cron string `db:"task_cron" json:"task_cron"`
	User_id   int    `db:"user_id" json:"user_id"`
	Url       string `db:"url" json:"url"`
	Id        int    `db:"id" json:"id"`
	Timestamp string `db:"_timestamp" json:"timestamp"`
}
