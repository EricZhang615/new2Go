package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	var build func(left, right int) *TreeNode
	build = func(left, right int) *TreeNode {
		if left == right {
			return nil
		}
		val, idx := nums[left], left
		for i := left; i < right; i++ {
			if nums[i] > val {
				val = nums[i]
				idx = i
			}
		}
		root := &TreeNode{Val: val}
		root.Left = build(left, idx)
		root.Right = build(idx+1, right)
		return root
	}
	return build(0, len(nums))
}

func main() {
	constructMaximumBinaryTree([]int{3, 2, 1})
}
