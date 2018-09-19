package leetcode

import "strconv"

/**
 * Definition for a binary tree node.
 *
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(t *TreeNode) string {
	result := ""
	if t == nil {
		return result
	}
	result += strconv.Itoa(t.Val)
	// 提升效率
	if t.Left == nil && t.Right == nil {
		return result
	}
	// 注意emit的条件
	if t.Left != nil || t.Right != nil {
		result += "(" + tree2str(t.Left) + ")"
	}
	if t.Right != nil {
		result += "(" + tree2str(t.Right) + ")"
	}
	return result
}
