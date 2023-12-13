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

func reverseList(head *ListNode) *ListNode {
	var i *ListNode
	i = nil
	j := head
	var tmp *ListNode
	for j != nil {
		tmp = j.Next
		j.Next = i
		i = j
		j = tmp
	}
	head = i
	return head
}
