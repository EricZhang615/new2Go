package main

type LinkedNode struct {
	Val  int
	Next *LinkedNode
}

type MyLinkedList struct {
	DummyNode *LinkedNode
	Size      int
}

func Constructor() MyLinkedList {
	newNode := &LinkedNode{0, nil}
	return MyLinkedList{newNode, 0}
}

func (this *MyLinkedList) Get(index int) int {
	if index >= this.Size || index < 0 {
		return -1
	}
	n := this.DummyNode.Next
	for i := 0; i < index; i++ {
		n = n.Next
	}
	return n.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.Size, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.Size {
		return
	}
	newNode := &LinkedNode{val, nil}
	n := this.DummyNode
	for i := -1; i < index-1; i++ {
		n = n.Next
	}
	newNode.Next = n.Next
	n.Next = newNode
	this.Size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.Size || index < 0 {
		return
	}
	n := this.DummyNode
	for i := 0; i < index; i++ {
		n = n.Next
	}
	n.Next = n.Next.Next
	this.Size--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
