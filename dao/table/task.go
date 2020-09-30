package table

import "time"

type Tasks struct {
	ID         int    `gorm:"primary_key;auto_increment"`
	TaskName   string `gorm:"type:varchar(255);not null"`
	TaskCron   string `gorm:"type:varchar(255);not null"`
	Url        string `gorm:"type:varchar(255);not null"`
	UserId     string `gorm:"type:varchar(255);not null"`
	ReqType    int    `gorm:"type:int;default 0;not null"`
	TaskStatus int    `gorm:"type:int;default 0;not null"`

	_Timestamp time.Time
}

func (p Tasks) TableName() string {
	return "tasks"
}
