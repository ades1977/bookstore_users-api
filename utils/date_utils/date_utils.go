package date_utils

import (
	"time"
)

const (
	apiDateLayout = "2006-01-02 15:04:05" //dibuat "2006-01-02 15:04:05"  sebagai default jam sistem
)

func GetNow() time.Time {
		return time.Now().UTC()
}

func GetLocalNow() time.Time {
	return time.Now()
}

func GetNowString() string{
	return  GetNow().Format(apiDateLayout)
}

func GetNowLocalString() string{
	return  GetLocalNow().Format(apiDateLayout)
}