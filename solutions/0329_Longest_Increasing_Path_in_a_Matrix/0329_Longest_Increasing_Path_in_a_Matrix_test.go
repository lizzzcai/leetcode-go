package leetcode

import (
	"fmt"
	"testing"
)

func TestLongestIncreasingPath(t *testing.T) {
	var tests = []struct {
		matrix [][]int
		want   int
	}{
		{[][]int{[]int{9, 9, 4}, []int{6, 6, 8}, []int{2, 1, 1}}, 4},
		{[][]int{[]int{3, 4, 5}, []int{3, 2, 6}, []int{2, 2, 1}}, 4},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.matrix)
		t.Run(testname, func(t *testing.T) {
			ans := longestIncreasingPath(tt.matrix)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
