package asukatime

import (
	"fmt"
	"testing"
	"time"
)

func Test_GetLastMonthRange(t *testing.T) {
	tr := GetLastMonthRange()
	fmt.Println(tr)
}

func Test_GetMonthRange(t *testing.T) {
	ti := time.Now()
	tr := GetMonthRange(ti)
	fmt.Println(tr)
}

func Test_GetDayRange(t *testing.T) {
	ti := time.Now()
	tr := GetDayRange(ti)
	fmt.Println(tr)
}
