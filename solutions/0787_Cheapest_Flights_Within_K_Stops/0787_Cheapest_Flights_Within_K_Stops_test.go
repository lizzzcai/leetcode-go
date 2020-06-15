package leetcode

import (
	"fmt"
	"testing"
)

func TestFindCheapestPrice(t *testing.T) {
	var tests = []struct {
		n           int
		flights     [][]int
		src, dst, k int
		want        int
	}{
		{3, [][]int{[]int{0, 1, 100}, []int{1, 2, 100}, []int{0, 2, 500}}, 0, 2, 1, 200},
		{3, [][]int{[]int{0, 1, 100}, []int{1, 2, 100}, []int{0, 2, 500}}, 0, 2, 0, 500},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %v, %v, %v, %v", tt.n, tt.flights, tt.src, tt.dst, tt.k)
		t.Run(testname, func(t *testing.T) {
			ans := findCheapestPrice(tt.n, tt.flights, tt.src, tt.dst, tt.k)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
