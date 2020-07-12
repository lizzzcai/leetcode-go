package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDiagonalSort(t *testing.T) {
	var tests = []struct {
		mat  [][]int
		want [][]int
	}{
		{[][]int{[]int{3, 3, 1, 1}, []int{2, 2, 1, 2}, []int{1, 1, 1, 2}}, [][]int{[]int{1, 1, 1, 1}, []int{1, 2, 2, 2}, []int{1, 2, 3, 3}}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.mat)
		t.Run(testname, func(*testing.T) {
			ans := diagonalSort(tt.mat)
			if reflect.DeepEqual(ans, tt.want) != true {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
