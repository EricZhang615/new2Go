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
	traversal(node.Right, ans)
	*ans = append(*ans, node.Val)
}

func postorderTraversal(root *TreeNode) []int {
	var ans []int
	traversal(root, &ans)
	return ans
}
