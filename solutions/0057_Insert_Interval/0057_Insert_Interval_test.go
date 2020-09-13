package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	var tests = []struct {
		intervals   [][]int
		newInterval []int
		want        [][]int
	}{
		{[][]int{{1, 3}, {6, 9}}, []int{2, 5}, [][]int{{1, 5}, {6, 9}}},
		{[][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}, [][]int{{1, 2}, {3, 10}, {12, 16}}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %v", tt.intervals, tt.newInterval)
		t.Run(testname, func(t *testing.T) {
			ans := insert(tt.intervals, tt.newInterval)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
