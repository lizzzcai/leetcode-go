package leetcode

import (
	"fmt"
	"testing"
)

func TestPathSum(t *testing.T) {
	node := &TreeNode{10, nil, nil}
	node.Left = &TreeNode{5, nil, nil}
	node.Right = &TreeNode{-3, nil, nil}

	node.Left.Left = &TreeNode{3, nil, nil}
	node.Left.Right = &TreeNode{2, nil, nil}

	node.Right.Right = &TreeNode{11, nil, nil}

	node.Left.Left.Left = &TreeNode{3, nil, nil}
	node.Left.Left.Right = &TreeNode{-2, nil, nil}

	node.Left.Right.Right = &TreeNode{1, nil, nil}

	var tests = []struct {
		root *TreeNode
		sum  int
		want int
	}{
		{node, 8, 3},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %v", tt.root, tt.sum)
		t.Run(testname, func(t *testing.T) {
			ans := pathSum(tt.root, tt.sum)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}

}
