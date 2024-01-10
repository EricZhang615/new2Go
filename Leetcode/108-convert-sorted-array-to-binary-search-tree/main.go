package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	var traversal func(nums []int, left int, right int) *TreeNode
	traversal = func(nums []int, left int, right int) *TreeNode {
		if left > right {
			return nil
		}
		mid := left + ((right - left) / 2)
		node := &TreeNode{Val: nums[mid]}
		node.Left = traversal(nums, left, mid-1)
		node.Right = traversal(nums, mid+1, right)
		return node
	}
	return traversal(nums, 0, len(nums)-1)
}
