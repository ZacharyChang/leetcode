// Given a linked list, swap every two adjacent nodes and return its head.

// For example,
// Given 1->2->3->4, you should return the list as 2->1->4->3.

// Your algorithm should use only constant space. You may not modify the values in the list, only nodes itself can be changed.

package leetcode

/**
 * Definition for singly-linked list.
 *
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	// nil node
	if head == nil {
		return nil
	}
	// one node
	if head.Next == nil {
		return head
	}
	second := head.Next
	third := head.Next.Next
	second.Next = head
	head.Next = swapPairs(third)
	return second
}
