package myutils

import (
	"errors"
	"strconv"
	"time"
)

var Loc, _ = time.LoadLocation("Asia/Shanghai")

var DateFormats = []string{
	"2006-01-02",
	"2006-01-2",
	"2006-1-02",
	"2006-1-2",
	"2006/01/02",
	"2006/1/02",
	"2006/01/2",
	"2006/1/2",
	"01-02-06",
	"1-2-06",
	"1-02-06",
	"01-2-06",
	"01/02/06",
	"1/2/06",
	"01/2/06",
	"01/02/06",
	"1/02/06",
}

// ConvertDateToTime 将日期字符串转换成time.Time格式
func ConvertDateToTime(dateStr string) (time.Time, error) {
	if dateStr != "" {
		var t1 time.Time
		var err error
		formats2 := []string{
			" 15", " 15:04", " 15:04:05", "",
		}
		for _, format := range DateFormats {
			for _, v := range formats2 {
				t1, err = time.ParseInLocation(format+v, dateStr, Loc)
				if err == nil {
					return t1, nil
				}
			}
		}
	}
	return time.Time{}, errors.New("-1")
}

// ConvertDateToUnix 日期字符串格式化为时间戳
func ConvertDateToUnix(dateStr string, category ...int8) int64 {
	if dateStr != "" {
		if len(category) >= 1 {
			if category[0] == 1 {
				dateStr = dateStr + " 00:00:00"
			}
			if category[0] == 2 {
				dateStr = dateStr + " 23:59:59"
			}
		}
		t1, err := ConvertDateToTime(dateStr)
		if err == nil {
			return int64(t1.Unix())
		}
	}

	return 0
}

// ConvertTimeToDate 时间戳转换成日期格式
func ConvertTimeToDate(dateInt int64, category ...int8) string {
	if len(category) == 1 {
		if category[0] == 1 {
			return time.Unix(dateInt, 0).In(Loc).Format("2006-01-02 15:04:05")
		}
		if category[0] == 2 {
			return time.Unix(dateInt, 0).In(Loc).Format("2006-01-02 15:04")
		}
	}
	return time.Unix(dateInt, 0).In(Loc).Format("2006-01-02")
}

// GetNowUnixTime 获取当前日期时间戳
func GetNowUnixTime() int64 {
	return time.Now().In(Loc).Unix()
}

// GetNowDateStr 获取当前日期
func GetNowDateStr(cate ...int8) string {
	t := time.Now().In(Loc).Unix()
	if len(cate) == 1 {
		return ConvertTimeToDate(t, cate...)
	}
	return ConvertTimeToDate(t)
}

// GetUploadPath 获取上传文件目录地址
func GetUploadPath(disk ...string) string {
	now := time.Now()
	if len(disk) > 0 && disk[0] != "" {
		return "storage/" + strconv.Itoa(now.Year()) + "/" + strconv.Itoa(int(now.Month())) + strconv.Itoa(now.Day()) + "/"
	}
	return "storage/download/" + strconv.Itoa(now.Year()) + "/" + strconv.Itoa(int(now.Month())) + strconv.Itoa(now.Day()) + "/"
}
