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

func sumOfLeftLeaves(root *TreeNode) int {
	result := 0
	if root == nil {
		return result
	}
	if root.Left != nil {
		if root.Left.Left == nil && root.Left.Right == nil {
			result += root.Left.Val
		}
	}
	result += sumOfLeftLeaves(root.Left)
	result += sumOfLeftLeaves(root.Right)
	return result
}
