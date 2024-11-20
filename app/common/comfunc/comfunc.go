package comfunc

import (
	"reflect"
	"time"
)

// 获取接口里指定的元素值
func Column(slice interface{}, fieldName string) []interface{} {
	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Kind() != reflect.Slice {
		panic("Provided value is not a slice")
	}

	result := make([]interface{}, sliceValue.Len())
	for i := 0; i < sliceValue.Len(); i++ {
		item := sliceValue.Index(i)
		value := GetField(item.Interface(), fieldName)
		result[i] = value
	}

	return result
}

func GetField(item interface{}, fieldName string) interface{} {
	r := reflect.ValueOf(item)
	f := reflect.Indirect(r).FieldByName(fieldName)
	return f.Interface()
}

// CurrentCalendar 获取当前日历
func CurrentCalendar() (int, time.Month, []int) {
	// 获取当前时间
	now := time.Now()
	year, month, _ := now.Date()

	// 获取该月的第一天和最后一天
	firstDay := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	lastDay := firstDay.AddDate(0, 1, -1)

	// 生成日历
	days := make([]int, 0)
	for day := firstDay; day.Before(lastDay.AddDate(0, 0, 1)); day = day.AddDate(0, 0, 1) {
		days = append(days, day.Day())
	}
	return year, month, days
}
