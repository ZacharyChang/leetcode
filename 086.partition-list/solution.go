package leetcode

import "fmt"

/**
 * Definition for singly-linked list.
 *
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// Array
// TODO: improve with Two Pointers
func partition(head *ListNode, x int) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	front := make([]int, 0)
	tail := make([]int, 0)

	for cur := head; cur != nil; cur = cur.Next {
		if cur.Val < x {
			front = append(front, cur.Val)
		} else {
			tail = append(tail, cur.Val)
		}
	}
	fmt.Println(front)

	nums := append(front, tail...)
	fmt.Println(nums)

	var cur, res *ListNode
	for _, num := range nums {
		if cur == nil {
			res = &ListNode{num, nil}
			cur = res
		} else {
			cur.Next = &ListNode{num, nil}
			cur = cur.Next
		}
	}
	return res
}
