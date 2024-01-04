package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	var compare func(left *TreeNode, right *TreeNode) bool
	compare = func(left *TreeNode, right *TreeNode) bool {
		if (left == nil && right != nil) || (left != nil && right == nil) {
			return false
		} else if left == nil && right == nil {
			return true
		} else if left.Val != right.Val {
			return false
		}
		return compare(left.Left, right.Right) && compare(left.Right, right.Left)
	}
	return compare(root.Left, root.Right)
}
