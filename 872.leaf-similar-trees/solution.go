package leetcode

import (
	"reflect"
)

/**
 * Definition for a binary tree node.
 *
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	return reflect.DeepEqual(leafArray(root1), leafArray(root2))
}

func leafArray(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{
			root.Val,
		}
	}
	result = append(result, leafArray(root.Left)...)
	result = append(result, leafArray(root.Right)...)
	return result
}
