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
	start := &ListNode{}
	start.Next = head
	pre := start
	for pre.Next != nil {
		cur := pre.Next
		for cur.Next != nil && cur.Next.Val == cur.Val {
			cur = cur.Next
		}
		if cur != pre.Next {
			pre.Next = cur.Next
		} else {
			pre = pre.Next
		}
	}
	return start.Next
}
