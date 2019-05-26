package leetcode

/**
 * Definition for singly-linked list.
 *
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// TODO: Two Pointers
func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	cur, count := head, 1
	var pre *ListNode
	for cur.Next != nil {
		pre = cur
		cur = cur.Next
		count++
	}
	if k >= count {
		return rotateRight(head, k%count)
	}
	pre.Next = nil
	cur.Next = head
	return rotateRight(cur, k-1)
}
