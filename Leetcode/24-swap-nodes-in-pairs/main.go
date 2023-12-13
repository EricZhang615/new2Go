package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	var i, j, tmp *ListNode
	i = head
	for i != nil && i.Next != nil {
		j = i.Next
		i.Next = j.Next
		j.Next = i
		if i == head {
			head = j
		} else {
			tmp.Next = j
		}
		tmp = i
		i = i.Next
	}
	return head
}
