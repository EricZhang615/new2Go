package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	var traversal func(root *TreeNode, count int) bool
	traversal = func(root *TreeNode, count int) bool {
		if root.Left == nil && root.Right == nil && count == 0 {
			return true
		}
		if root.Left == nil && root.Right == nil {
			return false
		}
		if root.Left != nil {
			if traversal(root.Left, count-root.Left.Val) {
				return true
			}
		}
		if root.Right != nil {
			if traversal(root.Right, count-root.Right.Val) {
				return true
			}
		}
		return false
	}
	if root == nil {
		return false
	}
	return traversal(root, targetSum-root.Val)
}
