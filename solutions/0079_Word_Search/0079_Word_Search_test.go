package leetcode

import (
	"fmt"
	"testing"
)

func TestExist(t *testing.T) {
	var tests = []struct {
		board [][]byte
		word  string
		want  bool
	}{
		{board: [][]byte{[]byte{'A', 'B', 'C', 'E'}, []byte{'S', 'F', 'C', 'S'}, []byte{'A', 'D', 'E', 'E'}}, word: "ABCCED", want: true},
		{board: [][]byte{[]byte{'A', 'B', 'C', 'E'}, []byte{'S', 'F', 'C', 'S'}, []byte{'A', 'D', 'E', 'E'}}, word: "SEE", want: true},
		{board: [][]byte{[]byte{'A', 'B', 'C', 'E'}, []byte{'S', 'F', 'C', 'S'}, []byte{'A', 'D', 'E', 'E'}}, word: "ABCB", want: false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %s", tt.board, tt.word)
		t.Run(testname, func(t *testing.T) {
			ans := exist(tt.board, tt.word)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
