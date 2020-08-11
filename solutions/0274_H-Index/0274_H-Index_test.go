package leetcode

import (
	"fmt"
	"testing"
)

func TestHIndex(t *testing.T) {
	var tests = []struct {
		citations []int
		want      int
	}{
		{[]int{3, 0, 6, 1, 5}, 3},
		{[]int{3, 0, 6, 1, 5, 7, 8, 9, 11}, 5},
		{[]int{}, 0},
		{[]int{100}, 1},
		{[]int{0}, 0},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.citations)
		t.Run(testname, func(t *testing.T) {
			ans := hIndex(tt.citations)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
