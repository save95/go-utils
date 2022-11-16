package timeutil

import (
	"time"
)

// GetMonday 获取本周周一的日期
func GetMonday() time.Time {
	now := time.Now()

	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}

	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).
		AddDate(0, 0, offset)
}

//func StatDayWith(startDay, endDay string) ([]uint, error) {
//	r, err := computeStatDay(startDay, endDay)
//	if nil != err {
//		monday := GetMonday()
//		start, err := strconv.Atoi(monday.Format("20060102"))
//		if nil != err {
//			return nil, err
//		}
//
//		sunday := monday.Add(6 * 24 * time.Hour)
//		end, err := strconv.Atoi(sunday.Format("20060102"))
//		if nil != err {
//			return nil, err
//		}
//
//		return []uint{uint(start), uint(end)}, nil
//	}
//
//	return r, nil
//}
//
//func computeStatDay(startDay, endDay string) ([]uint, error) {
//	startAt, err := time.ParseInLocation("2006-01-02", startDay, time.Local)
//	if nil != err {
//		return nil, err
//	}
//	start, err := strconv.Atoi(startAt.Format("20060102"))
//	if nil != err {
//		return nil, err
//	}
//
//	endAt, err := time.ParseInLocation("2006-01-02", endDay, time.Local)
//	if nil != err {
//		return nil, err
//	}
//	end, err := strconv.Atoi(endAt.Format("20060102"))
//	if nil != err {
//		return nil, err
//	}
//
//	if start > end {
//		return nil, xerror.New("开始时间不能在结束时间之后")
//	}
//
//	return []uint{
//		uint(start), uint(end),
//	}, nil
//}
