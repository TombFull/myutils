package myutils

import (
	"fmt"
	"testing"
)

func TestGetNowTime(t *testing.T) {
	fmt.Println(GetNowUnixTime())
}
