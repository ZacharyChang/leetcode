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
	if root == nil {
		return result
	}
	len := length(root)
	avg := len / k
	last := len % k
	tmp := root
	result[0] = root
	for i := 1; i < k; i++ {
		result[i] = result[i-1]
		if last != 0 && i <= last {
			tmp = result[i]
			result[i] = result[i].Next
		}
		for j := 0; j < avg && result[i] != nil; j++ {
			tmp = result[i]
			result[i] = result[i].Next
		}
		tmp.Next = nil
	}
	if len < k {
		for i := len; i < k; i++ {
			result[i] = nil
		}
	}
	return result
}

func length(root *ListNode) int {
	res := 0
	for root != nil {
		root = root.Next
		res++
	}
	return res
}
