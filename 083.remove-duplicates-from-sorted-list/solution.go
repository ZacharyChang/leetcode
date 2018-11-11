package leetcode

/**
 * Definition for singly-linked list.
 *
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	i := head
	j := head.Next
	for j != nil {
		if j.Val != i.Val {
			i.Next = j
			i = j
		}
		j = j.Next
	}
	i.Next = nil
	return head
}
