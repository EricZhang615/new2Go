package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	maxCnt := math.MinInt
	cnt := 0
	var res []int
	var pre *TreeNode
	var traversal func(root *TreeNode)
	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		traversal(root.Left)
		if pre == nil {
			cnt = 1
		} else {
			if root.Val == pre.Val {
				cnt++

			} else {
				cnt = 1
			}
		}
		if cnt == maxCnt {
			res = append(res, root.Val)
		} else if cnt > maxCnt {
			maxCnt = cnt
			res = append([]int{}, root.Val)
		}
		pre = root
		traversal(root.Right)
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	traversal(root)
	return res
}
