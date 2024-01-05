package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
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

		val := preorder[0]
		preorder = preorder[1:]
		root := &TreeNode{Val: val}

		// 根据 val 在中序遍历的位置，将中序遍历划分成左右两颗子树
		// 由于我们每次都从先序遍历的头取元素，所以要先遍历左子树再遍历右子树
		inorderRootIndex := idxMap[val]
		root.Left = build(inorderLeft, inorderRootIndex-1)
		root.Right = build(inorderRootIndex+1, inorderRight)
		return root
	}
	return build(0, len(inorder)-1)
}
