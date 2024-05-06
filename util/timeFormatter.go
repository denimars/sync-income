package util

import (
	"time"
)

func FormatToUtc(date string) string {
	format := "2006-01-02 15:04:05"
	location, _ := time.LoadLocation("Asia/Makassar")
	time_, err := time.ParseInLocation(format, date, location)
	if err != nil {
		panic(err)
	}
	utcTime := time_.UTC()
	return utcTime.Format(format)
}
