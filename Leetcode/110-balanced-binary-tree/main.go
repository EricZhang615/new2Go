package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	var getHeight func(root *TreeNode) int
	getHeight = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := getHeight(root.Left)
		if left == -1 {
			return -1
		}
		right := getHeight(root.Right)
		if right == -1 {
			return -1
		}
		if math.Abs(float64(left-right)) > 1.0 {
			return -1
		}
		return max(left, right) + 1
	}
	return getHeight(root) != -1
}
