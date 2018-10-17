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

func sumNumbers(root *TreeNode) int {
	result := 0
	helper(root, &result, 0)
	return result
}

func helper(root *TreeNode, sum *int, n int) {
	if root == nil {
		return
	}
	n = n*10 + root.Val
	if root.Left == nil && root.Right == nil {
		// use pointer to improve performance
		*sum += n
		return
	}
	helper(root.Left, sum, n)
	helper(root.Right, sum, n)
}
