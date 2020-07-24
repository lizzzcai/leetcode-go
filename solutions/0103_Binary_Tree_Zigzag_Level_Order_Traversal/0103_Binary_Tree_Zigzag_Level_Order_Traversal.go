package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	left_to_right := true
	var node *TreeNode
	for len(stack) > 0 {
		next_layer := make([]*TreeNode, 0)
		curr_res := make([]int, 0)
		for i := len(stack); i > 0; i-- {
			node, stack = stack[0], stack[1:]

			if left_to_right {
				curr_res = append(curr_res, node.Val)
			} else {
				curr_res = append([]int{node.Val}, curr_res...)
			}

			if node.Left != nil {
				next_layer = append(next_layer, node.Left)
			}

			if node.Right != nil {
				next_layer = append(next_layer, node.Right)
			}
		}
		res = append(res, curr_res)
		stack = next_layer
		left_to_right = !left_to_right
	}
	return res
}
