package leetcode

import (
	"fmt"
	"testing"
)

func TestWordBreak(t *testing.T) {
	var tests = []struct {
		s        string
		wordDict []string
		want     []string
	}{
		{"catsanddog", []string{"cat", "cats", "and", "sand", "dog"}, []string{"cats and dog", "cat sand dog"}},
		{"pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"}, []string{"pine apple pen apple", "pineapple pen apple", "pine applepen apple"}},
		{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}, []string{}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %v", tt.s, tt.wordDict)
		t.Run(testname, func(t *testing.T) {
			ans := wordBreak(tt.s, tt.wordDict)
			if !equal(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func equal(s, t []string) bool {
	if len(s) != len(t) {
		return false
	}
	hmap := make(map[string]bool)
	for _, x := range s {
		hmap[x] = true
	}

	for _, x := range t {
		_, ok := hmap[x]
		if !ok {
			return false
		}
	}

	return true
}
