package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestKthSmallest(t *testing.T) {
	var tests = []struct {
		matrix [][]int
		k      int
		want   int
	}{
		{[][]int{[]int{1, 5, 9}, []int{10, 11, 13}, []int{12, 13, 15}}, 8, 13},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v %v", tt.matrix, tt.k)
		t.Run(testname, func(t *testing.T) {
			ans := kthSmallest(tt.matrix, tt.k)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
