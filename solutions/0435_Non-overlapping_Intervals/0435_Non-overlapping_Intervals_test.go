package leetcode

import (
	"fmt"
	"testing"
)

func TestEraseOverlapIntervals(t *testing.T) {
	var tests = []struct {
		intervals [][]int
		want      int
	}{
		{[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}, 1},
		{[][]int{{1, 2}, {1, 2}, {1, 2}}, 2},
		{[][]int{{1, 2}, {2, 3}}, 0},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.intervals)
		t.Run(testname, func(t *testing.T) {
			ans := eraseOverlapIntervals(tt.intervals)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
