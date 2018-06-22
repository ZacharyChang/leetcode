// Given a linked list, remove the nth node from the end of list and return its head.

// For example,

//    Given linked list: 1->2->3->4->5, and n = 2.

//    After removing the second node from the end, the linked list becomes 1->2->3->5.
// Note:
// Given n will always be valid.
// Try to do this in one pass.

package leetcode

/**
 * Definition for singly-linked list.
 *
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	array := make([]*ListNode, 0)
	cur := head
	for cur != nil {
		array = append(array, cur)
		cur = cur.Next
	}
	pos := len(array) - n
	if pos == 0 {
		head = head.Next
	} else {
		if pos+1 < len(array) {
			array[pos-1].Next = array[pos+1]
		} else {
			array[pos-1].Next = nil
		}
	}
	return head
}
