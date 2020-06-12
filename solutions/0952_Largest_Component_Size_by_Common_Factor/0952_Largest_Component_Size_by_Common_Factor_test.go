package leetcode

import (
	"fmt"
	"testing"
)

func TestLargestComponentSize(t *testing.T) {
	var tests = []struct {
		A    []int
		want int
	}{
		{[]int{4, 6, 15, 35}, 4},
		{[]int{20, 50, 9, 63}, 2},
		{[]int{2, 3, 6, 7, 4, 12, 21, 39}, 8},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.A)
		t.Run(testname, func(t *testing.T) {
			ans := largestComponentSize(tt.A)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
