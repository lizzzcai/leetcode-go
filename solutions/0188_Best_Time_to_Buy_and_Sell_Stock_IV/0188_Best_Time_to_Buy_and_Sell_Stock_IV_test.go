package leetcode

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	var tests = []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{2, 4, 1}, 2, 2},
		{[]int{3, 2, 6, 5, 0, 3}, 2, 7},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v %v", tt.k, tt.nums)
		t.Run(testname, func(t *testing.T) {
			ans := maxProfit(tt.k, tt.nums)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
