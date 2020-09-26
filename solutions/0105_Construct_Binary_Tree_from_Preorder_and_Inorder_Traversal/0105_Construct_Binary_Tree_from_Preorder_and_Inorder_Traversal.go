package leetcode

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	m := make(map[int]int)
	for i, x := range inorder {
		m[x] = i
	}

	return helper(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1, m)
}

func helper(preorder, inorder []int, pi, pj, ii, ij int, m map[int]int) *TreeNode {
	if ii > ij {
		return nil
	}

	if ii == ij {
		return &TreeNode{inorder[ii], nil, nil}
	}

	root := &TreeNode{preorder[pi], nil, nil} // first val in preorder is the root
	idx, _ := m[root.Val]                     // index of root in inorder
	lsize := idx - ii                         // size of the left sub-tree
	root.Left = helper(preorder, inorder, pi+1, pi+lsize, ii, idx-1, m)
	root.Right = helper(preorder, inorder, pi+lsize+1, pj, idx+1, ij, m)
	return root
}
