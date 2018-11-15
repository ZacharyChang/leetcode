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
	slow := head
	quick := head
	for quick != nil && quick.Next != nil {
		slow = slow.Next
		quick = quick.Next.Next
	}
	return slow
}
