package leetcode

import (
	"fmt"
	"testing"
)

func TestCanVisitAllRooms(t *testing.T) {
	var tests = []struct {
		rooms [][]int
		want  bool
	}{
		{[][]int{{1}, {2}, {3}, {}}, true},
		{[][]int{{1, 3}, {3, 0, 1}, {2}, {0}}, false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.rooms)
		t.Run(testname, func(t *testing.T) {
			ans := canVisitAllRooms(tt.rooms)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
