package leetcode

import (
	"fmt"
	"testing"
)

func TestMinStickers(t *testing.T) {
	var tests = []struct {
		stickers []string
		target   string
		want     int
	}{
		{[]string{"with", "example", "science"}, "thehat", 3},
		{[]string{"notice", "possible"}, "basicbasic", -1},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %v", tt.stickers, tt.target)
		t.Run(testname, func(t *testing.T) {
			ans := minStickers(tt.stickers, tt.target)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
