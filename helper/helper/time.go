package helper

import "time"

func Time2String(t time.Time) string {
	timeTemplate := "2006-01-02 15:04:05"
	return time.Unix(t.Unix(), 0).Format(timeTemplate)
}

func String2Date(str string) (t time.Time, err error) {
	return time.ParseInLocation("2006-01-02", str, time.Local)
}

func String2Time(str string) (t time.Time, err error) {
	return time.ParseInLocation("2006-01-02 15:04:05", str, time.Local)
}

func UTCString2Time(str string) (t time.Time, err error) {
	return time.ParseInLocation("2006-01-02T15:04:05+08:00", str, time.Local)
}

func Time2UtcString(t time.Time) string {
	timeTemplate := "2006-01-02T15:04:05Z"
	return time.Unix(t.Unix(), 0).Format(timeTemplate)
}
