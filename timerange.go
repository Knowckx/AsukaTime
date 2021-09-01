package asukatime

import (
	"fmt"
	"time"
)

var (
	BackOffset = -1 * time.Millisecond
)

type TimeRange struct {
	Start time.Time
	End   time.Time
}

func (s TimeRange) String() string {
	out := fmt.Sprintf("[%s, %s]", s.Start, s.End)
	return out
}

func (s *TimeRange) ToUnix() *TimeRangeUnix {
	out := &TimeRangeUnix{}
	out.Start = s.Start.Unix()
	out.End = s.End.Unix()
	return out
}

type TimeRangeUnix struct {
	Start int64
	End   int64
}

// GetDayStart 1945-10-10 12:12:12 --> 1945-10-10 00:00:00
func GetDayStart(t time.Time) time.Time {
	tiStart := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, timeZone)
	return tiStart
}

// GetDayRange 1945-10-10 12:12:12 --> [1945-10-10 00:00:00, 1945-10-10 23:59:59]
func GetDayRange(t time.Time) *TimeRange {
	tiStart := GetDayStart(t)
	tiEnd := tiStart.AddDate(0, 0, 1).Add(BackOffset)
	tr := &TimeRange{
		tiStart, tiEnd,
	}
	return tr
}

// GetMonthRange 1945-10-10 12:12:12 --> [1945-09-01 00:00:00, 1945-09-30 23:59:59]
func GetYesterdayRange() *TimeRange {
	dayStart := GetDayStart(time.Now())
	ti := dayStart.AddDate(0, 0, -1)
	return GetDayRange(ti)
}

// GetWeekRange 1945-10-10 12:12:12 --> [1945-10-08 00:00:00, 1945-10-14 23:59:59]
func GetWeekRange(in time.Time) *TimeRange {
	weekDay := int(in.Weekday())
	dayFix := 1 // monday is the first day of week
	weekDay -= dayFix
	dayStart := GetDayStart(in)
	tiStart := dayStart.AddDate(0, 0, -1*weekDay)
	tiEnd := tiStart.AddDate(0, 0, 7).Add(BackOffset)
	tr := &TimeRange{
		tiStart, tiEnd,
	}
	return tr
}

// GetMonthRange 1945-10-10 12:12:12 --> [1945-10-01 00:00:00, 1945-10-30 23:59:59]
func GetMonthRange(ti time.Time) *TimeRange {
	today := GetDayStart(ti)
	dayOffset := ti.Day()*-1 + 1
	start := today.AddDate(0, 0, dayOffset)
	end := start.AddDate(0, 1, 0).Add(BackOffset)
	tr := TimeRange{
		start, end,
	}
	return &tr
}

// GetMonthRange 1945-10-10 12:12:12 --> [1945-09-01 00:00:00, 1945-09-30 23:59:59]
func GetLastMonthRange() *TimeRange {
	ti := time.Now().AddDate(0, -1, 0)
	return GetMonthRange(ti)
}

// GetYearRange 1945-10-10 12:12:12 --> [1945-01-01 00:00:00, 1945-12-31 23:59:59]
func GetYearRange(ti time.Time) *TimeRange {
	st := time.Date(ti.Year(), time.January, 1, 0, 0, 0, 0, ti.Location())
	end := st.AddDate(1, 0, 0).Add(BackOffset)
	tr := TimeRange{
		st, end,
	}
	return &tr
}

func GenMonthList(from, end int) []time.Time {
	outTime := []time.Time{}
	now := time.Now()
	for i := from; i <= end; i++ {
		ti := time.Date(now.Year(), time.Month(i), now.Day(), 0, 0, 0, 0, time.UTC)
		outTime = append(outTime, ti)

	}
	return outTime
}
