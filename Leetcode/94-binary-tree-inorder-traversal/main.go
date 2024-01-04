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

func inorderTraversal2(root *TreeNode) (res []int) {
	var stack []*TreeNode
	cur := root
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, cur.Val)
			cur = cur.Right
		}
	}
	return res
}
