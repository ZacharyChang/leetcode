package leetcode

/**
 * Definition for singly-linked list.
 *
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(root *ListNode, k int) []*ListNode {
	result := make([]*ListNode, k)

	len := func(cur *ListNode) int {
		res := 0
		for cur != nil {
			cur = cur.Next
			res++
		}
		return res
	}(root)

	div, mod := len/k, len%k
	cur := root
	for i := 0; i < k; i++ {
		result[i] = cur
		num := div
		if i < mod {
			num++
		}
		for j := 0; j < num-1; j++ {
			cur = cur.Next
		}
		if cur != nil {
			cur, cur.Next = cur.Next, nil
		}
	}
	return result
}
