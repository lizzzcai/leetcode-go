package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

type item struct {
	Node  *TreeNode
	Index int
}

func widthOfBinaryTree(root *TreeNode) int {
	ans := 0
	if root == nil {
		return ans
	}

	curr_layer := make([]*item, 0)
	next_layer := make([]*item, 0)
	curr_layer = append(curr_layer, &item{root, 0})
	var N, curr_width int
	var left, right *item
	for len(curr_layer) > 0 {
		// update the current width of current layer
		N = len(curr_layer)
		left, right = curr_layer[0], curr_layer[N-1]
		curr_width = right.Index - left.Index + 1
		if curr_width > ans {
			ans = curr_width
		}

		for _, pop := range curr_layer {
			node := pop.Node
			index := pop.Index

			if node.Left != nil {
				next_layer = append(next_layer, &item{node.Left, index*2 + 1})
			}

			if node.Right != nil {
				next_layer = append(next_layer, &item{node.Right, index*2 + 2})
			}
		}
		curr_layer = next_layer
		next_layer = []*item{}
	}
	return ans
}
