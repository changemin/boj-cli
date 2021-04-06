package utils

import "time"

func GetCurrentDate() string {
	dateTime := time.Now()
	return dateTime.Format("2006-01-02")
}
