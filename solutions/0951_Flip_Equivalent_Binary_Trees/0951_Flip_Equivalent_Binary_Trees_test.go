package leetcode

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func deserialize(s string) *TreeNode {
	if len(s) == 0 {
		return nil
	}
	s_split := strings.Split(strings.Trim(s, "[]"), ",")
	nodes := make([]*TreeNode, len(s))
	for i, val := range s_split {
		if val == "null" {
			nodes[i] = nil
		} else {
			int_val, err := strconv.Atoi(val)
			if err != nil {
				nodes[i] = &TreeNode{int_val, nil, nil}
			}
		}
	}
	// reverse nodes
	kids := make([]*TreeNode, len(s))
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		kids[i], kids[j] = nodes[j], nodes[i]
	}

	root := kids[len(kids)-1]
	kids = kids[:len(kids)-1]

	for _, node := range nodes {
		if node != nil {
			if kids[len(kids)-1] != nil {
				node.Left = kids[len(kids)-1]
				kids = kids[:len(kids)-1]
			}
			if kids[len(kids)-1] != nil {
				node.Right = kids[len(kids)-1]
				kids = kids[:len(kids)-1]
			}
		}
	}
	return root
}

func TestFlipEquiv(t *testing.T) {
	var tests = []struct {
		root1, root2 string
		want         bool
	}{
		{"[1,2,3,4,5,6,null,null,null,7,8]", "[1,3,2,null,6,4,5,null,null,null,null,8,7]", true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %v", tt.root1, tt.root2)
		t.Run(testname, func(t *testing.T) {
			ans := flipEquiv(deserialize(tt.root1), deserialize(tt.root2))
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
