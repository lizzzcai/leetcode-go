package leetcode

import (
	"fmt"
	"testing"
)

func TestLargestTimeFromDigits(t *testing.T) {
	var tests = []struct {
		A    []int
		want string
	}{
		{[]int{1, 2, 3, 4}, "23:41"},
		{[]int{5, 5, 5, 5}, ""},
		{[]int{0, 0, 0, 0}, "00:00"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.A)
		t.Run(testname, func(t *testing.T) {
			ans := largestTimeFromDigits(tt.A)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
