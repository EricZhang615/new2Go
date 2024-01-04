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

// 递归法

func preorderTraversal3(root *TreeNode) (res []int) {
	var stack []*TreeNode
	if root == nil {
		return res
	}
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return res
}
