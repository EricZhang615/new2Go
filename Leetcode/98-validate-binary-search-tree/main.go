package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var pre *TreeNode
	var valid func(root *TreeNode) bool
	valid = func(root *TreeNode) bool {
		if root == nil {
			return true
		}
		left := valid(root.Left)
		if pre != nil && pre.Val >= root.Val {
			return false
		}
		pre = root
		right := valid(root.Right)
		return left && right
	}
	return valid(root)
}
