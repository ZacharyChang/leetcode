package leetcode

// recursive
// better understanding but worse performance
func merge_2(a *ListNode, b *ListNode) *ListNode {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	var start *ListNode
	if a.Val < b.Val {
		start = a
		start.Next = merge_2(a.Next, b)
	} else {
		start = b
		start.Next = merge_2(a, b.Next)
	}
	return start
}

func sortList_2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := middle(head)
	next := mid.Next
	mid.Next = nil

	// care: must sort before merge
	res := merge_2(sortList(head), sortList(next))
	return res
}
