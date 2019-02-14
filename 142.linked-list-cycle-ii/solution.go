package leetcode

/**
 * Definition for singly-linked list.
 *
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// Hash Table
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	cur := head
	nodeMap := make(map[*ListNode]bool, 0)
	for cur != nil {
		if nodeMap[cur] {
			return cur
		}
		nodeMap[cur] = true
		cur = cur.Next
	}
	return nil
}
