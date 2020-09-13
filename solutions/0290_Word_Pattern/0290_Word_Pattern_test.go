package leetcode

import (
	"fmt"
	"testing"
)

func TestWordPattern(t *testing.T) {
	var tests = []struct {
		pattern string
		str     string
		want    bool
	}{
		{"abba", "dog cat cat dog", true},
		{"abba", "dog cat cat fish", false},
		{"abba", "dog dog dog dog", false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %v", tt.pattern, tt.str)
		t.Run(testname, func(t *testing.T) {
			ans := wordPattern(tt.pattern, tt.str)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
