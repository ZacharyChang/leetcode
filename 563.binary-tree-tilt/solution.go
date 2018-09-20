package leetcode

/**
 * Definition for a binary tree node.
 *
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := sum(root.Left)
	right := sum(root.Right)
	return abs(left-right) + findTilt(root.Left) + findTilt(root.Right)
}

func sum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return root.Val + sum(root.Left) + sum(root.Right)
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
