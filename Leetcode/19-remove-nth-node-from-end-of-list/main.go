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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyNode := &ListNode{0, head}
	var i *ListNode
	i = dummyNode
	l := []*ListNode{dummyNode}
	for i.Next != nil {
		i = i.Next
		l = append(l, i)
	}
	l[len(l)-n-1].Next = l[len(l)-n].Next
	return dummyNode.Next
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	cur := head
	prev := dummyHead
	i := 1
	for cur != nil {
		cur = cur.Next
		if i > n {
			prev = prev.Next
		}
		i++
	}
	prev.Next = prev.Next.Next
	return dummyHead.Next
}
