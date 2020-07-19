package vojo

type TasksHistory struct {
	Task_id     int    `db:"task_id"`
	Exec_time   string `db:"exec_time"`
	Exec_result string `db:"exec_result"`
	Exec_code   int    `db:"exec_code"`
	Id          int    `db:"id"`
	Timestamp   string `db:"_timestamp"`
}
