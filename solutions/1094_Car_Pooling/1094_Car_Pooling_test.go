package leetcode

import (
	"fmt"
	"testing"
)

func TestCarPooling(t *testing.T) {
	var tests = []struct {
		trips    [][]int
		capacity int
		want     bool
	}{
		{[][]int{{3, 2, 7}, {3, 7, 9}, {8, 3, 9}}, 11, true},
		{[][]int{{2, 1, 5}, {3, 5, 7}}, 3, true},
		{[][]int{{2, 1, 5}, {3, 3, 7}}, 4, false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %v", tt.trips, tt.capacity)
		t.Run(testname, func(t *testing.T) {
			ans := carPooling(tt.trips, tt.capacity)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
