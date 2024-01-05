package main

import "slices"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func traversal(inorder []int, postorder []int) *TreeNode {
	// 遍历序列空 返回空节点
	if len(postorder) == 0 {
		return nil
	}
	// 后序末尾为中间节点 初始化
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	// 序列长度为1 是叶子节点 可以去掉，因为递归到空的时候上面已经包含了
	if len(postorder) == 1 {
		return root
	}
	// 用中间节点切中序序列
	index := slices.Index(inorder, root.Val)
	leftInorder := inorder[:index]
	rightInorder := inorder[index+1:]
	// 用左中序和右中序切后序序列
	postorder = postorder[:len(postorder)-1]
	leftPostorder := postorder[:len(leftInorder)]
	rightPostorder := postorder[len(leftInorder):]
	// 递归
	root.Left = traversal(leftInorder, leftPostorder)
	root.Right = traversal(rightInorder, rightPostorder)

	return root
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	return traversal(inorder, postorder)
}

func buildTree2(inorder []int, postorder []int) *TreeNode {
	idxMap := map[int]int{}
	for i, v := range inorder {
		idxMap[v] = i
	}
	var build func(int, int) *TreeNode
	build = func(inorderLeft, inorderRight int) *TreeNode {
		// 无剩余节点
		if inorderLeft > inorderRight {
			return nil
		}

		// 后序遍历的末尾元素即为当前子树的根节点
		val := postorder[len(postorder)-1]
		postorder = postorder[:len(postorder)-1]
		root := &TreeNode{Val: val}

		// 根据 val 在中序遍历的位置，将中序遍历划分成左右两颗子树
		// 由于我们每次都从后序遍历的末尾取元素，所以要先遍历右子树再遍历左子树
		inorderRootIndex := idxMap[val]
		root.Right = build(inorderRootIndex+1, inorderRight)
		root.Left = build(inorderLeft, inorderRootIndex-1)
		return root
	}
	return build(0, len(inorder)-1)
}

func main() {
	buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
}
