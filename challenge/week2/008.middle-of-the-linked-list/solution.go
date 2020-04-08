package leetcode

/**
 * Definition for singly-linked list.
 *
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	first := head
	second := head
	for second != nil && second.Next != nil {
		first = first.Next
		second = second.Next.Next
	}
	return first
}
