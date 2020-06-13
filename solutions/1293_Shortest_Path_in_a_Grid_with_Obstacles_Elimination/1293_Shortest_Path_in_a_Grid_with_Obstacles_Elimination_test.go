package leetcode

import (
	"fmt"
	"testing"
)

func TestShortestPath(t *testing.T) {
	var tests = []struct {
		grid [][]int
		k    int
		want int
	}{
		{[][]int{[]int{0, 0, 0}, []int{1, 1, 0}, []int{0, 0, 0}, []int{0, 1, 1}, []int{0, 0, 0}}, 1, 6},
		{[][]int{[]int{0, 1, 1}, []int{1, 1, 1}, []int{1, 0, 0}}, 1, -1},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %v", tt.grid, tt.k)
		t.Run(testname, func(t *testing.T) {
			ans := shortestPath(tt.grid, tt.k)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
