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

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return helper(root, 1)
}

func helper(root *TreeNode, depth int) int {
	if root.Left == nil && root.Right == nil {
		return depth
	}
	depth++
	if root.Left == nil {
		return helper(root.Right, depth)
	}
	if root.Right == nil {
		return helper(root.Left, depth)
	}
	return min(helper(root.Left, depth), helper(root.Right, depth))
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
