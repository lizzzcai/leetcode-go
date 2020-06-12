package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	var tests = []struct {
		nums []int
		want []int
	}{
		{[]int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
		{[]int{2, 0, 1}, []int{0, 1, 2}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.nums)
		t.Run(testname, func(t *testing.T) {
			tmp := tt.nums
			sortColors(tmp)
			if !reflect.DeepEqual(tmp, tt.want) {
				t.Errorf("got %v, want %v", tmp, tt.want)
			}

		})
	}
}
