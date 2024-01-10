package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	var traversal func(root *TreeNode)
	pre := 0
	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		traversal(root.Right)
		root.Val += pre
		pre = root.Val
		traversal(root.Left)
		return
	}
	traversal(root)
	return root
}
