package leetcode

import (
	"fmt"
	"testing"
)

func TestWidthOfBinaryTree(t *testing.T) {

	node1 := &TreeNode{1, nil, nil}
	node1.Left = &TreeNode{3, nil, nil}
	node1.Right = &TreeNode{2, nil, nil}

	node1.Left.Left = &TreeNode{5, nil, nil}
	node1.Left.Right = &TreeNode{3, nil, nil}
	node1.Right.Right = &TreeNode{9, nil, nil}

	node2 := &TreeNode{1, nil, nil}
	node2.Left = &TreeNode{3, nil, nil}

	node2.Left.Left = &TreeNode{5, nil, nil}
	node2.Left.Right = &TreeNode{3, nil, nil}

	var tests = []struct {
		root *TreeNode
		want int
	}{
		{node1, 4},
		{node2, 2},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.root)
		t.Run(testname, func(t *testing.T) {
			ans := widthOfBinaryTree(tt.root)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
