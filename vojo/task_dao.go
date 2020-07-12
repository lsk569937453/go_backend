package vojo

type TasksDao struct {
	Task_name string `db:"task_name"`
	Task_cron string `db:"task_cron"`
	User_id   int    `db:"user_id"`
	Url       string `db:"url"`
	Id        int    `db:"id"`
	Timestamp string `db:"_timestamp"`
}
