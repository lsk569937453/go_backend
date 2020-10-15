package util

import "time"

const TimeFormatFirst = "2006-01-02-15-04-05-000"
const TimeFormatSecond = "2006-01-02 15:04:05.000"

func GetCurrentTime() string {
	currentTime := time.Now().Format(TimeFormatFirst)
	return currentTime
}
func FormatTime(needTime time.Time) string {
	return needTime.Format(TimeFormatSecond)

}
