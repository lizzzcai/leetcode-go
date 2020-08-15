package leetcode

import (
	"fmt"
	"testing"
)

func TestOrangesRotting(t *testing.T) {
	var tests = []struct {
		grid [][]int
		want int
	}{
		{[][]int{[]int{2, 1, 1}, []int{1, 1, 0}, []int{0, 1, 1}}, 4},
		{[][]int{[]int{2, 1, 1}, []int{0, 1, 1}, []int{1, 0, 1}}, -1},
		{[][]int{[]int{0, 2}}, 0},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.grid)
		t.Run(testname, func(t *testing.T) {
			ans := orangesRotting(tt.grid)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
