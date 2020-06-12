package leetcode

import (
	"fmt"
	"testing"
)

func TestLongestMountain(t *testing.T) {
	var tests = []struct {
		A    []int
		want int
	}{
		{[]int{2, 1, 4, 7, 3, 2, 5}, 5},
		{[]int{2, 2, 2}, 0},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.A)
		t.Run(testname, func(t *testing.T) {
			ans := longestMountain(tt.A)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
