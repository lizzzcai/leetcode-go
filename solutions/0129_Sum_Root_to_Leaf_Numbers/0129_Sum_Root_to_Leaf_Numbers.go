package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var res = 0
	dfs(root, 0, &res)
	return res
}

func dfs(node *TreeNode, curr int, res *int) {
	curr = curr*10 + node.Val
	if node.Left == nil && node.Right == nil {
		*res += curr
		return
	}

	if node.Left != nil {
		dfs(node.Left, curr, res)
	}
	if node.Right != nil {
		dfs(node.Right, curr, res)
	}
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type item struct {
	Node *TreeNode
	Sum  int
}

func sumNumbers_1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var queue = make([]item, 0)
	queue = append(queue, item{root, root.Val})

	var res = 0
	for len(queue) > 0 {
		pop := queue[0]
		queue = queue[1:]

		if pop.Node.Left == nil && pop.Node.Right == nil {
			res += pop.Sum
		}

		if pop.Node.Left != nil {
			queue = append(queue, item{pop.Node.Left, pop.Sum*10 + pop.Node.Left.Val})
		}

		if pop.Node.Right != nil {
			queue = append(queue, item{pop.Node.Right, pop.Sum*10 + pop.Node.Right.Val})
		}
	}

	return res
}
