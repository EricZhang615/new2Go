package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	var res [][]int
	var path []int
	var traversal func(root *TreeNode, count int)
	traversal = func(root *TreeNode, count int) {
		if root.Left == nil && root.Right == nil && count == 0 {
			res = append(res, append([]int{}, path...))
			return
		}
		if root.Left != nil {
			path = append(path, root.Left.Val)
			traversal(root.Left, count-root.Left.Val)
			path = path[:len(path)-1]
		}
		if root.Right != nil {
			path = append(path, root.Right.Val)
			traversal(root.Right, count-root.Right.Val)
			path = path[:len(path)-1]
		}
		return
	}
	if root == nil {
		return res
	}
	path = append(path, root.Val)
	traversal(root, targetSum-root.Val)
	return res
}
