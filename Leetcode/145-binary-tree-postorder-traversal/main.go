package main

import "slices"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func traversal(node *TreeNode, ans *[]int) {
	if node == nil {
		return
	}
	traversal(node.Left, ans)
	traversal(node.Right, ans)
	*ans = append(*ans, node.Val)
}

func postorderTraversal(root *TreeNode) []int {
	var ans []int
	traversal(root, &ans)
	return ans
}

func postorderTraversal2(root *TreeNode) (res []int) {
	var stack []*TreeNode
	if root == nil {
		return res
	}
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	slices.Reverse(res)
	return res
}
