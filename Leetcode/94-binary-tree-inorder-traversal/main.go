package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func traversal(node *TreeNode, ans *[]int) {
	if node == nil {
		return
	}
	traversal(node.Left, ans)
	*ans = append(*ans, node.Val)
	traversal(node.Right, ans)
}

func inorderTraversal(root *TreeNode) []int {
	var ans []int
	traversal(root, &ans)
	return ans
}
