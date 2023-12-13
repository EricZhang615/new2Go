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

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	dummyNode1 := &ListNode{0, headA}
	dummyNode2 := &ListNode{0, headB}
	i, j := dummyNode1, dummyNode2
	l1, l2 := 0, 0
	for i.Next != nil {
		i = i.Next
		l1++
	}
	for j.Next != nil {
		j = j.Next
		l2++
	}
	i, j = dummyNode1, dummyNode2
	if l1 > l2 {
		for k := 0; k < l1-l2; k++ {
			i = i.Next
		}
	} else {
		for k := 0; k < l2-l1; k++ {
			j = j.Next
		}
	}
	for i.Next != nil && j.Next != nil {
		if i.Next == j.Next {
			return i.Next
		}
		i = i.Next
		j = j.Next
	}
	return nil
}
