package leetcode

import (
	"fmt"
	"testing"
)

func TestFirstMissingPositive(t *testing.T) {
	var tests = []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 0}, 3},
		{[]int{3, 4, -1, 1}, 2},
		{[]int{7, 8, 9, 11, 12}, 1},
		{[]int{}, 1},
		{[]int{1}, 2},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.nums)
		t.Run(testname, func(t *testing.T) {
			ans := firstMissingPositive(tt.nums)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
