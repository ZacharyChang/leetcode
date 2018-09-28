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

func largestValues(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	array := []*TreeNode{
		root,
	}
	for len(array) > 0 {
		newArray := make([]*TreeNode, 0)
		max := array[0].Val
		for _, v := range array {
			if v.Val > max {
				max = v.Val
			}
			if v.Left != nil {
				newArray = append(newArray, v.Left)
			}
			if v.Right != nil {
				newArray = append(newArray, v.Right)
			}
		}
		result = append(result, max)
		array = newArray
	}
	return result
}
