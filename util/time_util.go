package util

import "time"

const TimeFormat = "2006-01-02-15:04:05.000"

func GetCurrentTime() string {
	currentTime := time.Now().Format(TimeFormat)
	return currentTime
}
