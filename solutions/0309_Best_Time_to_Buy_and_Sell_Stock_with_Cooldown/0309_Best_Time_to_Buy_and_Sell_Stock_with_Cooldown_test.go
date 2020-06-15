package leetcode

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	var tests = []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 0, 2}, 3},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.nums)
		t.Run(testname, func(t *testing.T) {
			ans := maxProfit(tt.nums)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
