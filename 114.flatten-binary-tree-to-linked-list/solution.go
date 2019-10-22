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

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

	if root.Left == nil {
		return
	}
	left := root.Left
	for left.Right != nil {
		left = left.Right
	}
	left.Right = root.Right
	root.Right = root.Left
	root.Left = nil
}
