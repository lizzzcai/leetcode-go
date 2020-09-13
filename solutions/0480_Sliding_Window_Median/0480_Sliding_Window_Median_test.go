package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMedianSlidingWindow(t *testing.T) {
	var tests = []struct {
		nums []int
		k    int
		want []float64
	}{
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []float64{1.00000, -1.00000, -1.00000, 3.00000, 5.00000, 6.00000}},
		{[]int{1, 2, 3, 4, 2, 3, 1, 4, 2}, 3, []float64{2.00000, 3.00000, 3.00000, 3.00000, 2.00000, 3.00000, 2.00000}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %v", tt.nums, tt.k)
		t.Run(testname, func(t *testing.T) {
			ans := medianSlidingWindow(tt.nums, tt.k)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
