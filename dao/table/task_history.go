package table

import "time"

type TaskExecHistory struct {
	ID         int    `gorm:"primary_key;auto_increment"`
	TaskId     int    `gorm:"type:int;default 0;not null"`
	ExecTime   string `gorm:"type:varchar(255);not null"`
	ExecResult string `gorm:"type:longText;not null"`
	ExecCode   int    `gorm:"type:int;default 0;not null"`
	TaskCron   string `gorm:"type:longText;not null"`
	TaskUrl    string `gorm:"type:varchar(255);not null"`

	Timestamp time.Time
}

func (p TaskExecHistory) TableName() string {
	return "task_exec_history"
}
