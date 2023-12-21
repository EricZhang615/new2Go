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
	*ans = append(*ans, node.Val)
	traversal(node.Left, ans)
	traversal(node.Right, ans)
}

func preorderTraversal(root *TreeNode) []int {
	var ans []int
	traversal(root, &ans)
	return ans
}

func preorderTraversal2(root *TreeNode) (res []int) {
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		traversal(node.Left)
		traversal(node.Right)
	}
	traversal(root)
	return res
}
