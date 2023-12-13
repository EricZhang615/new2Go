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

func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	n := dummyHead
	for n != nil && n.Next != nil {
		if n.Next.Val == val {
			n.Next = n.Next.Next
		} else {
			n = n.Next
		}
	}
	return dummyHead.Next
}
