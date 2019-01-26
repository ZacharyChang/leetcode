package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

type MyLinkedList struct {
	head   *ListNode
	length int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{
		nil,
		0,
	}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index >= this.length {
		return -1
	}
	cur := this.head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	return cur.Val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.length, val)
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.length {
		return
	}
	this.length++
	if index == 0 {
		this.head = &ListNode{
			val,
			this.head,
		}
		return
	}
	cur, tmp := this.head.Next, this.head
	for i := 1; i < index; i++ {
		tmp = cur
		cur = cur.Next
	}
	cur = &ListNode{
		val,
		cur,
	}
	tmp.Next = cur
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.length {
		return
	}
	this.length--
	if index == 0 {
		this.head = this.head.Next
		return
	}
	cur, tmp := this.head.Next, this.head
	for i := 1; i < index; i++ {
		tmp = cur
		cur = cur.Next
	}
	tmp.Next = cur.Next
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
