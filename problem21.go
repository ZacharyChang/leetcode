// Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

// Example:

// Input: 1->2->4, 1->3->4
// Output: 1->1->2->3->4->4

package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l2.Val >= l1.Val {
		temp := l1.Next
		l1.Next = mergeTwoLists(temp, l2)
		head = l1
	} else {
		temp := l2.Next
		l2.Next = mergeTwoLists(l1, temp)
		head = l2
	}
	return head
}
