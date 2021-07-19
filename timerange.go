package asukatime

import (
	"fmt"
	"time"
)

type TimeRange struct {
	Start time.Time
	End   time.Time
}

func (s TimeRange) String() string {
	out := fmt.Sprintf("[%s, %s]", s.Start, s.End)
	return out
}

// GetDayRange 1945-10-10 12:12:12 --> [1945-10-10 00:00:00, 1945-10-10 23:59:59]
func GetDayRange(t time.Time) *TimeRange {
	tiStart := GetDayStart(t)
	tiEnd := tiStart.AddDate(0, 0, 1).Add(-1 * time.Second)
	tr := &TimeRange{
		tiStart, tiEnd,
	}
	return tr
}

// GetDayStart 1945-10-10 12:12:12 --> 1945-10-10 00:00:00
func GetDayStart(t time.Time) time.Time {
	tiStart := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return tiStart
}

// GetMonthRange 1945-10-10 12:12:12 --> [1945-10-01 00:00:00, 1945-10-30 23:59:59]
func GetMonthRange(ti time.Time) *TimeRange {
	today := GetDayStart(ti)
	dayOffset := ti.Day()*-1 + 1
	start := today.AddDate(0, 0, dayOffset)
	end := start.AddDate(0, 1, 0).Add(-1 * time.Second)
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
