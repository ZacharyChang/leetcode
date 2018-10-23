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

var result []int

func preorderTraversal(root *TreeNode) []int {
	result = make([]int, 0)
	helper(root)
	return result
}

func helper(root *TreeNode) {
	if root == nil {
		return
	}
	result = append(result, root.Val)
	helper(root.Left)
	helper(root.Right)
}
