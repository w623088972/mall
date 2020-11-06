package util

import (
	"fmt"
	"time"
)

const timeFmt = "2006-01-02 15:04:05"

//Len return character number in string(not byte number)
func StrLen(src string) int {
	i := 0
	for range src {
		i++
	}
	return i
}

func ParseStrTime(timeStr string) (time.Time, error) {
	t, err := time.Parse(timeFmt, timeStr)
	return t, err
}

func ToRFC3339(now time.Time) string {
	return now.Format(timeFmt)
}

func StrNow() string {
	return time.Now().Format(timeFmt)
}

func ConvertRFC3339TimeFormat(t string) (string, error) {
	tmp, err := time.Parse(time.RFC3339, t)
	if err != nil {
		return t, nil
	}
	return tmp.Format(timeFmt), nil
}

//获取相差时间（秒）
func GetSecondDiffer(startTime, endTime string) int64 {
	var second int64
	t1, err := time.ParseInLocation(timeFmt, startTime, time.Local)
	t2, err := time.ParseInLocation(timeFmt, endTime, time.Local)
	if err == nil && t1.Before(t2) {
		diff := t2.Unix() - t1.Unix()
		second = diff
		return second
	} else {
		return second
	}
}

func GetTimeUnixToTimeStr(unix int) string {
	tm := time.Unix(int64(unix), 0)
	return tm.Format(timeFmt)
}

//时间字符串转时间戳(string —> int64)
func StrToUnix(date string) int64 {
	local, _ := time.LoadLocation("Local")                   //重要：获取时区
	theTime, _ := time.ParseInLocation(timeFmt, date, local) //使用模板在对应时区转化为time.time类型
	unix := theTime.Unix()                                   //转化为时间戳，类型是int64

	return unix
}

//根据起止时间获取之间的连续时间
func GetDateFromRange(startTime, endTime string) []string {
	startUnix := StrToUnix(startTime)
	endUnix := StrToUnix(endTime)

	//计算日期段内有多少天
	days := (endUnix-startUnix)/86400 + 1;

	//保存每天日期
	var date []string
	for i := 0; i < int(days); i++ {
		date = append(date, time.Unix(startUnix+int64(86400*i), 0).Format("2006-01-02"))
	}

	return date
}

//Yesterday 返回昨天 起始、截止时间 xxxx-xx-xx 00:00:00 xxxx-xx-xx 23:59:59
func Yesterday() (string, string) {
	y, m, d := time.Now().AddDate(0, 0, -1).Date()
	start := fmt.Sprintf("%d-%02d-%02d 00:00:00", y, m, d)
	end := fmt.Sprintf("%d-%02d-%02d 23:59:59", y, m, d)

	return start, end
}
