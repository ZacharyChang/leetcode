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

func isCousins(root *TreeNode, x int, y int) bool {
	return depth(root, x) == depth(root, y) && !sameParent(root, x, y)
}

func depth(root *TreeNode, x int) int {
	if root == nil {
		return -1
	}
	if root.Val == x {
		return 0
	}
	if left := depth(root.Left, x); left != -1 {
		return left + 1
	}
	if right := depth(root.Right, x); right != -1 {
		return right + 1
	}
	return -1
}

func sameParent(root *TreeNode, x int, y int) bool {
	if root == nil {
		return false
	}
	if root.Left != nil && root.Right != nil {
		if root.Left.Val == x && root.Right.Val == y {
			return true
		}
		if root.Left.Val == y && root.Right.Val == x {
			return true
		}
	}
	return sameParent(root.Left, x, y) || sameParent(root.Right, x, y)
}
