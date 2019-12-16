package util

import (
	"fmt"
	"time"
)

const MINUTES_OF_DAY = 24 * 60

func GetTimePeriod(date time.Time, periodCount int) []string {
	if MINUTES_OF_DAY%periodCount != 0 {
		fmt.Println("MINUTES_OF_DAY % periodCount !=0")
		panic("MINUTES_OF_DAY % periodCount !=0")
	}

	timeStr := date.Format("2006-01-02")
	standardDate, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr+" 00:00:00", time.UTC)
	if err != nil {
		fmt.Println(err)
	}

	//定义Slice
	var timePeriodArray []string

	minutes := MINUTES_OF_DAY / periodCount
	startTime := standardDate
	for i := 0; i < periodCount; i++ {
		endTime := startTime.Add(time.Duration(minutes) * time.Minute)
		startTimeStr := startTime.Format("2006-01-02T 15:04:05Z")
		endTimeStr := endTime.Format("2006-01-02T 15:04:05Z")
		timePeriodArray = append(timePeriodArray, startTimeStr+","+endTimeStr)
		startTime = endTime
	}
	return timePeriodArray
}

func GetYesterdayZeroTime() time.Time {
	timeStr := time.Now().Format("2006-01-02")
	today, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr+" 00:00:00", time.UTC)

	yesterday := today.AddDate(0, 0, -1)
	return yesterday
}
