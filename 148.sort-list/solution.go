package leetcode

/**
 * Definition for singly-linked list.
 *
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := middle(head)
	next := mid.Next
	mid.Next = nil
	// care: must sort before merge
	return merge(sortList(head), sortList(next))
}

// loop
func merge(a *ListNode, b *ListNode) *ListNode {
	start := &ListNode{
		-1,
		nil,
	}
	cur := start
	for a != nil && b != nil {
		if a.Val < b.Val {
			cur.Next = a
			cur = cur.Next
			a = a.Next
		} else {
			cur.Next = b
			cur = cur.Next
			b = b.Next
		}
		if a == nil {
			cur.Next = b
		}
		if b == nil {
			cur.Next = a
		}
	}
	return start.Next
}

func middle(head *ListNode) *ListNode {
	slow := head
	quick := head
	pre := head
	for quick != nil && quick.Next != nil {
		pre = slow
		slow = slow.Next
		quick = quick.Next.Next
	}
	return pre
}
