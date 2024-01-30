package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minCameraCover(root *TreeNode) int {
	res := 0
	var traversal func(root *TreeNode) int
	traversal = func(root *TreeNode) int {
		if root == nil {
			return 2
		}
		left := traversal(root.Left)
		right := traversal(root.Right)
		if left == 2 && right == 2 {
			return 0
		}
		if left == 0 || right == 0 {
			res++
			return 1
		} else {
			return 2
		}
	}
	if traversal(root) == 0 {
		res++
	}
	return res
}
