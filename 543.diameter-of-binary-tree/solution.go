package leetcode

import (
	"sort"
)

/**
 * Definition for a binary tree node.
 *
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// recursive
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 0
	}
	return max(depth(root.Left)+depth(root.Right), diameterOfBinaryTree(root.Left), diameterOfBinaryTree(root.Right))
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(depth(root.Left), depth(root.Right)) + 1
}

func max(nums ...int) int {
	sort.Ints(nums)
	return nums[len(nums)-1]
}
