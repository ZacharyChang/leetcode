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

var count int

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	count = 0
	helper(root)
	return count
}

func helper(root *TreeNode) {
	count++
	if root.Left != nil {
		helper(root.Left)
	}
	if root.Right != nil {
		helper(root.Right)
	}
}
