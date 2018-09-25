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

func findBottomLeftValue(root *TreeNode) int {
	v, _ := findBottomLeftLevel(root, 0)
	return v
}

func findBottomLeftLevel(root *TreeNode, level int) (int, int) {
	if root == nil {
		return 0, 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val, level
	}
	leftValue, leftLevel := findBottomLeftLevel(root.Left, level+1)
	rightValue, rightLevel := findBottomLeftLevel(root.Right, level+1)
	if leftLevel > level && leftLevel >= rightLevel {
		return leftValue, leftLevel
	}
	return rightValue, rightLevel
}
