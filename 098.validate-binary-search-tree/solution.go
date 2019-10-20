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

func isValidBST(root *TreeNode) bool {
	arr := traversal(root)
	for i := 1; i < len(arr); i++ {
		if arr[i] <= arr[i-1] {
			return false
		}
	}
	return true
}

func traversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{
		root.Val,
	}
	if root.Left == nil && root.Right == nil {
		return res
	}
	res = append(traversal(root.Left), res...)
	res = append(res, traversal(root.Right)...)
	return res
}
