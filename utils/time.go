package utils

import (
	"github.com/lijianjunljj/gmfcommon/config"
	"time"
)

//DateTimeFormat 年月日时分秒时间格式
const DateTimeFormat = config.DateTimeFormat

//DateTime 当前时间
func DateTime() string {
	return time.Now().Format(DateTimeFormat)
}

//TimeUnix 获取当前的时间Unix时间戳
func TimeUnix() int64 {
	return time.Now().Unix()
}

//TimeMilliUnix 获取当前的时间毫秒级时间戳
func TimeMilliUnix() int64 {
	return time.Now().UnixNano() / 1000000
}

//TimeMilliUnixAdd 获取当前的时间毫秒级时间戳
func TimeMilliUnixAdd(d string) int64 {
	duration, _ := time.ParseDuration(d)
	return time.Now().Add(duration).UnixNano() / 1000000
}
func DurationToSeconds(d string) float64 {
	duration, _ := time.ParseDuration(d)
	return duration.Seconds()
}

//TimeNanoUnix 获取当前的时间纳秒级时间戳
func TimeNanoUnix() int64 {
	return time.Now().UnixNano()
}

//DateTime2MilliUnix 时间字符串转毫秒时间戳
func DateTime2MilliUnix(datetime string) int64 {
	dt, _ := time.ParseInLocation(datetime, datetime, time.Local)
	return dt.UnixNano() / 1000000
}

//MilliUnix2DateTime 毫秒时间戳转时间字符串
func MilliUnix2DateTime(millisecond int64) string {
	return time.Unix(millisecond, 0).Format(DateTimeFormat)
}

// GetCurrentTimestamp 获取当天的时间范围
//Time类型 2022-05-19 00:00:00 +0800 CST 2022-05-19 23:59:59 +0800 CST
func GetCurrentTimestamp() (beginTime, endTime time.Time) {
	t := time.Now()
	timeStr := t.Format("2006-01-02")
	beginTime, _ = time.ParseInLocation("2006-01-02", timeStr, time.Local)
	endTimeTmp := beginTime.Unix() + 86399
	endTime = time.Unix(endTimeTmp, 0)
	return beginTime, endTime
}
