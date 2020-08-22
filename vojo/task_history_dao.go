package vojo

type TasksHistory struct {
	Task_id     int    `db:"task_id" json:"task_id"`
	Exec_time   string `db:"exec_time" json:"exec_time"`
	Exec_result string `db:"exec_result" json:"exec_result"`
	Exec_code   int    `db:"exec_code" json:"exec_code"`
	Id          int    `db:"id" json:"id"`
	Timestamp   string `db:"_timestamp" json:"_timestamp"`
}
