package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestZigzagLevelOrder(t *testing.T) {
	node1 := &TreeNode{3, nil, nil}
	node1.Left = &TreeNode{9, nil, nil}
	node1.Right = &TreeNode{20, nil, nil}
	node1.Right.Left = &TreeNode{15, nil, nil}
	node1.Right.Right = &TreeNode{7, nil, nil}

	var tests = []struct {
		root *TreeNode
		want [][]int
	}{
		{node1, [][]int{[]int{3}, []int{20, 9}, []int{15, 7}}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.root)
		t.Run(testname, func(t *testing.T) {
			ans := zigzagLevelOrder(tt.root)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
