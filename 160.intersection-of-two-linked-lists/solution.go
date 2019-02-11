package leetcode

/**
 * Definition for singly-linked list.
 *
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// two pointers
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	ptrA, ptrB := headA, headB
	for ptrA != ptrB {
		if ptrA == nil {
			ptrA = headB
		} else {
			ptrA = ptrA.Next
		}
		if ptrB == nil {
			ptrB = headA
		} else {
			ptrB = ptrB.Next
		}
	}
	return ptrA
}
