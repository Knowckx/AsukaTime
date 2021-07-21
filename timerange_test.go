package asukatime

import (
	"fmt"
	"testing"
	"time"
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
