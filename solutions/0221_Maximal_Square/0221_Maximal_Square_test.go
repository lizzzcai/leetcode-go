package leetcode

import (
	"fmt"
	"testing"
)

func TestMaximalSquare(t *testing.T) {
	var tests = []struct {
		matrix [][]byte
		want   int
	}{
		{[][]byte{[]byte{'1', '0', '1', '0', '0'}, []byte{'1', '0', '1', '1', '1'}, []byte{'1', '1', '1', '1', '1'}, []byte{'1', '0', '0', '1', '0'}}, 4},
		{[][]byte{[]byte{}}, 0},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.matrix)
		t.Run(testname, func(t *testing.T) {
			ans := maximalSquare(tt.matrix)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}

		})
	}
}
