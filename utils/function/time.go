package function

/*
 * @Description:
 * @version: v1.0.0
 * @Author: shahao
 * @Date: 2022-03-16 19:45:39
 * @LastEditors: shahao
 * @LastEditTime: 2022-09-15 16:07:45
 */

import (
	"strconv"
	"time"
)

const (
	FmtDate              = "2006-01-02"
	FmtTime              = "15:04:05"
	FmtDateTime          = "2006-01-02 15:04:05"
	FmtDateTimeNoSeconds = "2006-01-02 15:04"
	FmtPayDateTime       = "20060102150405"
)

// 秒时间戳
func NowUnix() int64 {
	return time.Now().Unix()
}

// 秒时间戳转时间
func FromUnix(unix int64) time.Time {
	return time.Unix(unix, 0)
}

// 当前毫秒时间戳
func NowTimestamp() int64 {
	return Timestamp(time.Now())
}

// 毫秒时间戳
func Timestamp(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

// 毫秒时间戳转时间
func FromTimestamp(timestamp int64) time.Time {
	return time.Unix(0, timestamp*int64(time.Millisecond))
}

// 时间格式化
func Format(time time.Time, layout string) string {
	return time.Format(layout)
}

// 字符串时间转时间类型
func Parse(timeStr, layout string) (time.Time, error) {
	return time.Parse(layout, timeStr)
}

// return yyyyMMdd
func GetDay(time time.Time) int {
	ret, _ := strconv.Atoi(time.Format("20060102"))
	return ret
}

func WithTimeAsStartOfMin(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, t.Location())
}
func WithTimeAsEndOfMin(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 59, 0, t.Location())
}

// 返回指定时间当天的开始时间
func WithTimeAsStartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}
func WithTimeAsEndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
}

func NextYearTodayStartTimeTimestamp() int64 {
	now := time.Now().AddDate(+1, 0, 0)
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).UnixNano() / 1e6
}

func NextMonthStartTimeTimestamp(nextData int) int64 {
	now := time.Now().AddDate(0, nextData, 0)
	return time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location()).UnixNano()/1e6 - 1000
}

/**
 * 将时间格式换成 xx秒前，xx分钟前...
 * 规则：
 * 59秒--->刚刚
 * 1-59分钟--->x分钟前（23分钟前）
 * 1-24小时--->x小时前（5小时前）
 * 昨天--->昨天 hh:mm（昨天 16:15）
 * 前天--->前天 hh:mm（前天 16:15）
 * 前天以后--->mm-dd（2月18日）
 */
func PrettyTime(milliseconds int64) string {
	t := FromTimestamp(milliseconds)
	duration := (NowTimestamp() - milliseconds) / 1000
	if duration < 60 {
		return "刚刚"
	} else if duration < 3600 {
		return strconv.FormatInt(duration/60, 10) + "分钟前"
	} else if duration < 86400 {
		return strconv.FormatInt(duration/3600, 10) + "小时前"
	} else if Timestamp(WithTimeAsStartOfDay(time.Now().Add(-time.Hour*24))) <= milliseconds {
		return "昨天 " + Format(t, FmtTime)
	} else if Timestamp(WithTimeAsStartOfDay(time.Now().Add(-time.Hour*24*2))) <= milliseconds {
		return "前天 " + Format(t, FmtTime)
	} else {
		return Format(t, FmtDate)
	}
}

/*
*
获取一个时间点那一周，周一的开始时间和结束时间，毫秒，默认是 Sunday 开始到 Saturday 算 0,1,2,3,4,5,6
*/
func GetFirstDateOfWeek(t time.Time) (weekMonday int64) {
	// now := time.Now()
	offset := int(time.Monday - t.Weekday())
	if offset > 0 {
		offset = -6
	}

	weekStartDate := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	weekMonday = weekStartDate.UnixNano() / 1e6
	// weekSunday = weekMonday + 7*24*60*60*1000
	return
}

func GetFirstDateOfMonth(t time.Time) (monthFirstDay int64) {
	// now := time.Now()

	monthFirstDayDate := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local)
	monthFirstDay = monthFirstDayDate.UnixNano() / 1e6
	return
}

func GetRangeDateOfWeek(t time.Time, n int) (weekMonday, weekSunday int64) {
	weekMonday = GetFirstDateOfWeek(t)
	weekSunday = time.Unix(0, weekMonday*int64(time.Millisecond)).AddDate(0, 0, n).UnixNano()/1e6 - 1
	return
}

func GetRangeDateOfMonth(t time.Time, n int) (monthFirstDay, monthLastDay int64) {
	monthFirstDay = GetFirstDateOfMonth(t)
	monthLastDay = time.Unix(0, monthFirstDay*int64(time.Millisecond)).AddDate(0, n, 0).UnixNano()/1e6 - 1
	return
}

// 获取两个时间相差的天数，0表同一天，正数表t1>t2，负数表t1<t2
func GetDiffDays(t1, t2 time.Time) int {
	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.Local)

	return int(t1.Sub(t2).Hours() / 24)
}

// 获取t1和t2的相差天数，单位：秒，0表同一天，正数表t1>t2，负数表t1<t2
func GetDiffDaysBySecond(t1, t2 int64) int {
	time1 := time.Unix(t1, 0)
	time2 := time.Unix(t2, 0)

	// 调用上面的函数
	return GetDiffDays(time1, time2)
}
