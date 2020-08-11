package leetcode

import (
	"fmt"
	"testing"
)

func TestTitleToNumber(t *testing.T) {
	var tests = []struct {
		s    string
		want int
	}{
		{"A", 1},
		{"AB", 28},
		{"ZY", 701},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.s)
		t.Run(testname, func(t *testing.T) {
			ans := titleToNumber(tt.s)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
