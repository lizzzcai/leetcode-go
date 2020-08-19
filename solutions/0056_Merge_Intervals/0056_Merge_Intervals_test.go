package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	var tests = []struct {
		intervals [][]int
		want      [][]int
	}{
		{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.intervals)
		t.Run(testname, func(t *testing.T) {
			ans := merge(tt.intervals)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
