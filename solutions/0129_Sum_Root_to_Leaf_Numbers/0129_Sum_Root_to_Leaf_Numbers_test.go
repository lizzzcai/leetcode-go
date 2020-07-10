package leetcode

import (
	"fmt"
	"testing"
)

func TestSumNumbers(t *testing.T) {

	node1 := &TreeNode{1, nil, nil}
	node1.Left = &TreeNode{2, nil, nil}
	node1.Right = &TreeNode{3, nil, nil}

	node2 := &TreeNode{4, nil, nil}
	node2.Left = &TreeNode{9, nil, nil}
	node2.Right = &TreeNode{0, nil, nil}
	node2.Left.Left = &TreeNode{5, nil, nil}
	node2.Left.Right = &TreeNode{1, nil, nil}

	var tests = []struct {
		root *TreeNode
		want int
	}{
		{node1, 25},
		{node2, 1026},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.root)
		t.Run(testname, func(t *testing.T) {
			for _, fn := range []func(*TreeNode) int{sumNumbers, sumNumbers_1} {
				ans := fn(tt.root)
				if ans != tt.want {
					t.Errorf("got %v, want %v", ans, tt.want)
				}
			}

		})
	}
}
