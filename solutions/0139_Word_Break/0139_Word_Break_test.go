package leetcode

import (
	"fmt"
	"testing"
)

func TestWordBreak(t *testing.T) {
	var tests = []struct {
		s        string
		wordDict []string
		want     bool
	}{
		{"leetcode", []string{"leet", "code"}, true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %v", tt.s, tt.wordDict)
		t.Run(testname, func(t *testing.T) {
			ans := wordBreak(tt.s, tt.wordDict)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
