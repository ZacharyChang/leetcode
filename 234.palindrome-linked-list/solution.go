package leetcode

/**
 * Definition for singly-linked list.
 *
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// O(n) space
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	stack := make([]int, 0)
	mid := middle(head)
	for ; head != nil && head != mid; head = head.Next {
		stack = append(stack, head.Val)
	}
	for ; head != nil; head = head.Next {
		if stack[len(stack)-1] == head.Val {
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func middle(head *ListNode) *ListNode {
	slow := head
	quick := head
	for quick != nil && quick.Next != nil {
		slow = slow.Next
		quick = quick.Next.Next
	}
	return slow
}
