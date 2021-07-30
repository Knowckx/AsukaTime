package asukatime

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_GetDayRange_Now(t *testing.T) {
	ti := time.Now()
	tr := GetDayRange(ti)
	fmt.Println(tr)
}

func Test_GetWeekRange_Now(t *testing.T) {
	ti := time.Now()
	tr := GetWeekRange(ti)
	fmt.Println(tr)
}
func Test_GetLastMonthRange_Now(t *testing.T) {
	tr := GetLastMonthRange()
	fmt.Println(tr)
}

func Test_GetMonthRange_Now(t *testing.T) {
	ti := time.Now()
	tr := GetMonthRange(ti)
	fmt.Println(tr)
}

func Test_GetYearRange_Now(t *testing.T) {
	ti := time.Now()
	tr := GetYearRange(ti)
	fmt.Println(tr)
}

func Test_NewTime(t *testing.T) {
	in := "1945-10-10"
	// in := "1945-10-10 12:12:12"
	ti, err := NewTime(in)
	assert.Nil(t, err)
	fmt.Println(ti)
}

func Test_ToUnix(t *testing.T) {
	in := "2021-07-30"
	ti, err := NewTime(in)
	assert.Nil(t, err)
	ra := GetMonthRange(ti.UTC()).ToUnix()
	fmt.Println(ra)
}
