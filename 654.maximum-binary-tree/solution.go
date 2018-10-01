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

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	maxPos := findMaxPos(nums)
	return &TreeNode{
		nums[maxPos],
		constructMaximumBinaryTree(nums[:maxPos]),
		constructMaximumBinaryTree(nums[maxPos+1:]),
	}
}

func findMaxPos(nums []int) int {
	pos := 0
	for k, v := range nums {
		if v > nums[pos] {
			pos = k
		}
	}
	return pos
}
