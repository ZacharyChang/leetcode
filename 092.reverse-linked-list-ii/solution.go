package leetcode

/**
 * Definition for singly-linked list.
 *
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}
	count := 0
	cur := head
	var start, end, pre, after *ListNode
	for cur != nil {
		count++
		if count < m {
			pre = cur
		} else if count == m {
			start = cur
		}
		if count == n {
			end = cur
			after = end.Next
		}
		cur = cur.Next
	}
	if m == 1 {
		end.Next = nil
		reverseList(head)
		start.Next = after
		return end
	}
	end.Next = nil
	pre.Next = reverseList(start)
	start.Next = after
	return head
}

// Recursion
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	list := reverseList(next)
	next.Next = head
	head.Next = nil
	return list
}
