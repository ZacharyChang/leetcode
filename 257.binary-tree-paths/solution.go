package leetcode

import "fmt"

/**
 * Definition for a binary tree node.
 *
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var result []string

func binaryTreePaths(root *TreeNode) []string {
	result = make([]string, 0)
	helper(root, "")
	return result
}

func helper(root *TreeNode, curStr string) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		curStr = fmt.Sprintf("%s%d", curStr, root.Val)
		result = append(result, curStr)
	} else {
		curStr = fmt.Sprintf("%s%d->", curStr, root.Val)
		helper(root.Left, curStr)
		helper(root.Right, curStr)
	}
}
