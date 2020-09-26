package leetcode

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func serialize(root *TreeNode) string {
	arr := []string{}
	q := []*TreeNode{root}
	for len(q) > 0 {
		pop := q[0]
		q = q[1:]

		if pop == nil {
			arr = append(arr, "$")
		} else {
			arr = append(arr, strconv.Itoa(pop.Val))
			q = append(q, []*TreeNode{pop.Left, pop.Right}...)
		}
	}

	return strings.Join(arr, ",")
}

func TestBuildTree(t *testing.T) {
	var tests = []struct {
		preorder, inorder []int
		want              string
	}{
		{[]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}, "3,9,20,$,$,15,7,$,$,$,$"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %v", tt.preorder, tt.inorder)
		t.Run(testname, func(t *testing.T) {
			ans := serialize(buildTree(tt.preorder, tt.inorder))
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
