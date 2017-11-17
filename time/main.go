package main

import (
	"fmt"
	"reflect"
	"time"
)

const local = "Asia/Tokyo"
const dateFormat = "2006-01-02"
const timeFormat = "2006-01-02 15:04:05"

func init() {
	// set local time zone.
	loc, err := time.LoadLocation(local)
	if err != nil {
		loc = time.FixedZone(local, 9*60*60)
	}
	time.Local = loc
}

func main() {
	// get current time.
	t := time.Now()
	fmt.Println(t)
	// year
	fmt.Println(t.Year())
	// month
	fmt.Println(t.Month())
	// day
	fmt.Println(t.Day())
	// hour
	fmt.Println(t.Hour())
	// minute
	fmt.Println(t.Minute())
	// second
	fmt.Println(t.Second())
	// weekday
	fmt.Println(t.Weekday())

	createTimeStruct()
	timeToString()
	operationMonth()
}

func createTimeStruct() {
	// An error occurs if you do not specify for nanoseconds.
	t := time.Date(2017, 10, 01, 23, 59, 59, 0, time.Local)
	fmt.Println(t) // 2017-10-01 23:59:59 +0900 JST
}

func timeToString() {
	t := time.Now()
	// parse to date string.
	fmt.Println(t.Format(dateFormat), reflect.TypeOf(t.Format(dateFormat))) // 2017-11-15 string
	// parse to datetime string.
	fmt.Println(t.Format(timeFormat)) // 2017-11-15 10:44:32
}

func operationMonth() {
	t := time.Now()
	// get start of this month.
	start := t.AddDate(0, 0, -1*t.Day()+1).Format(dateFormat)
	fmt.Println(start)
	// get end of this month.
	end := t.AddDate(0, 1, -1*t.Day()).Format(dateFormat)
	fmt.Println(end)
}
