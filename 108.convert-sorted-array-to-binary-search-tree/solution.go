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

func sortedArrayToBST(nums []int) *TreeNode {
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, left int, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	return &TreeNode{
		nums[mid],
		helper(nums, left, mid-1),
		helper(nums, mid+1, right),
	}
}
