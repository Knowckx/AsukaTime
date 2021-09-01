package asukatime

import (
	"time"

	"github.com/pkg/errors"
)

const FormatLayoutDate = "2006-01-02"
const FormatLayoutTime = "2006-01-02 15:04:05"

var timeZone = time.UTC

func UseUtc() {
	timeZone = time.UTC
}

func UseLocal() {
	timeZone = time.Local
}

// NewTime accept time string like "2006-01-02" or "2006-01-02 15:04:05"
func NewTime(in string) (time.Time, error) {
	layout, err := GetTimeLayout(in)
	if err != nil {
		return time.Time{}, errors.WithStack(err)
	}
	out, err := time.ParseInLocation(layout, in, time.Local)
	if err != nil {
		return time.Time{}, errors.WithStack(err)
	}
	return out, err
}

// NewTime like NewTime.but return utc time
func NewTimeUTC(in string) (time.Time, error) {
	layout, err := GetTimeLayout(in)
	if err != nil {
		return time.Time{}, errors.WithStack(err)
	}
	out, err := time.ParseInLocation(layout, in, time.UTC)
	if err != nil {
		return time.Time{}, errors.WithStack(err)
	}
	return out, err
}

func GetTimeLayout(in string) (string, error) {
	layout := ""
	if len(in) == len(FormatLayoutTime) {
		layout = FormatLayoutTime
	} else if len(in) == len(FormatLayoutDate) {
		layout = FormatLayoutDate
	} else {
		return "", errors.Errorf("Check the time format:%s", in)
	}
	return layout, nil
}
