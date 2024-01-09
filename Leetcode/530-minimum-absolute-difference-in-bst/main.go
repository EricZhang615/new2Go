package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	var pre *TreeNode
	m := math.MaxInt
	var traversal func(root *TreeNode)
	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		traversal(root.Left)
		if pre != nil {
			m = min(root.Val-pre.Val, m)
		}
		pre = root
		traversal(root.Right)
	}
	traversal(root)
	return m
}
