package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindAllConcatenatedWordsInADict(t *testing.T) {
	var tests = []struct {
		words []string
		want  []string
	}{
		{[]string{"cat", "cats", "catsdogcats", "dog", "dogcatsdog", "hippopotamuses", "rat", "ratcatdogcat"}, []string{"catsdogcats", "dogcatsdog", "ratcatdogcat"}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.words)
		t.Run(testname, func(t *testing.T) {
			ans := findAllConcatenatedWordsInADict(tt.words)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
