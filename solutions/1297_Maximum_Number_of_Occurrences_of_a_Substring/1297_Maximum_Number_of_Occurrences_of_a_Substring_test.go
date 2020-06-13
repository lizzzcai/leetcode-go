package leetcode

import (
	"fmt"
	"testing"
)

func TestMaxFreq(t *testing.T) {
	var tests = []struct {
		s                            string
		maxLetters, minSize, maxSize int
		want                         int
	}{
		{"aababcaab", 2, 3, 4, 2},
		{"aaaa", 1, 3, 3, 2},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %v, %v, %v", tt.s, tt.maxLetters, tt.minSize, tt.maxSize)
		t.Run(testname, func(t *testing.T) {
			ans := maxFreq(tt.s, tt.maxLetters, tt.minSize, tt.maxSize)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}

}
