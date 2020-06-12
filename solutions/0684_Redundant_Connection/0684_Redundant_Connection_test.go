package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindRedundantConnection(t *testing.T) {
	var tests = []struct {
		edges [][]int
		want  []int
	}{
		{[][]int{[]int{1, 2}, []int{1, 3}, []int{2, 3}}, []int{2, 3}},
		{[][]int{[]int{1, 2}, []int{2, 3}, []int{3, 4}, []int{1, 4}, []int{1, 5}}, []int{1, 4}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.edges)
		t.Run(testname, func(t *testing.T) {
			ans := findRedundantConnection(tt.edges)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
