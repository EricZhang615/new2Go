package main

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	queue := list.New()
	if root != nil {
		queue.PushBack(root)
	}
	var res [][]int
	for queue.Len() > 0 {
		size := queue.Len()
		var l []int
		for i := 0; i < size; i++ {
			e := queue.Front()
			queue.Remove(e)
			node := e.Value.(*TreeNode)
			// node := queue.Remove(queue.Front()).(*TreeNode)    //出队列
			l = append(l, node.Val)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		res = append(res, l)
	}
	return res
}
