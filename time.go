package utils

import "time"

// 获取当天0点的时间戳
func DateTimeTodayStartTs() int64  {
	t := time.Now()
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Unix()
}

// 获取某天的开始时间戳
// date 格式必须是 2020-4-15
func DateTimeDayStartTs(date string) int64  {
	t, _ := time.ParseInLocation("2006-01-02", date, time.Local)
	return t.Unix()
}