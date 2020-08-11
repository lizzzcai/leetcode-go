package leetcode

import (
	"fmt"
	"testing"
)

func TestUniquePaths(t *testing.T) {
	var tests = []struct {
		m    int
		n    int
		want int
	}{
		{3, 2, 3},
		{7, 3, 28},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("m=%v, n=%v", tt.m, tt.n)
		t.Run(testname, func(t *testing.T) {
			ans := uniquePaths(tt.m, tt.n)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})

	}

}
