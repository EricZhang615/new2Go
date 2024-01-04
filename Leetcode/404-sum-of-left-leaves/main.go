package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		ans += root.Left.Val
	} else {
		ans += sumOfLeftLeaves(root.Left)
	}
	ans += sumOfLeftLeaves(root.Right)
	return ans
}
