package leetcode

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {
	return inorder(root, sum)
}

func inorder(node *TreeNode, sum int) int {
	var res int
	if node != nil {
		res += inorder(node.Left, sum)
		res += dfs(node, sum, sum)
		res += inorder(node.Right, sum)
	}

	return res
}

func dfs(node *TreeNode, curr, sum int) int {
	if node == nil {
		return 0
	}

	var res int
	if curr == node.Val {
		res += 1
	}

	if node.Left != nil {
		res += dfs(node.Left, curr-node.Val, sum)
	}
	if node.Right != nil {
		res += dfs(node.Right, curr-node.Val, sum)
	}

	return res
}
