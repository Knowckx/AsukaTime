package asukatime

import (
	"time"

	"github.com/pkg/errors"
)

const FormatLayoutDate = "2006-01-02"
const FormatLayoutTime = "2006-01-02 15:04:05"

// NewTime accept time string like "2006-01-02" or "2006-01-02 15:04:05"
func NewTime(in string) (time.Time, error) {
	layout := ""
	if len(in) == len(FormatLayoutTime) {
		layout = FormatLayoutTime
	} else if len(in) == len(FormatLayoutDate) {
		layout = FormatLayoutDate
	} else {
		return time.Time{}, errors.Errorf("Check the time format:%s", in)
	}
	out, err := time.ParseInLocation(layout, in, time.Local)
	if err != nil {
		return time.Time{}, errors.WithStack(err)
	}
	return out, err
}
