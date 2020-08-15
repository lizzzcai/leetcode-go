package leetcode

import (
	"fmt"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	var tests = []struct {
		s    string
		want int
	}{
		{"abccccdd", 7},
		{"abc", 1},
		{"aaa", 3},
		{"a", 1},
		{"", 0},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.s)
		t.Run(testname, func(t *testing.T) {
			ans := longestPalindrome(tt.s)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
