package leetcode

import (
	"fmt"
	"testing"
)

func TestConstrainedSubsetSum(t *testing.T) {
	var tests = []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{10, 2, -10, 5, 20}, 2, 37},
		{[]int{-1, -2, -3}, 1, -1},
		{[]int{10, -2, -10, -5, 20}, 2, 23},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("nums=%v, k=%v", tt.nums, tt.k)
		t.Run(testname, func(t *testing.T) {
			ans := constrainedSubsetSum(tt.nums, tt.k)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}

		})
	}
}
