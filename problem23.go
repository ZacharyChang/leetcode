// Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.

package leetcode

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	pos, minNode := findMinNode(lists)
	_, nextNode := findMinNode(append(lists[:pos], lists[pos+1:]...))

	minNode.Next = nextNode
	return minNode
}

func findMinNode(lists []*ListNode) (int, *ListNode) {
	if len(lists) == 0 {
		return 0, nil
	}
	if len(lists) == 1 {
		return 0, lists[0]
	}
	if lists[0] == nil {
		lists = lists[1:]
		return findMinNode(lists)
	}
	min := lists[0].Val
	pos := 0
	for i, list := range lists {
		if list == nil {
			continue
		}
		if list.Val < min {
			pos = i
			min = list.Val
		}
	}
	fmt.Print("pos: " + string(pos) + " , min: " + string(min))
	// printList(lists[pos])
	// printList(lists[pos])
	return pos, lists[pos]
}
