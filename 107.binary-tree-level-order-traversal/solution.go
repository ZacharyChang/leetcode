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

var result [][]int

func levelOrderBottom(root *TreeNode) [][]int {
	result = make([][]int, 0)
	if root == nil {
		return result
	}
	helper([]*TreeNode{
		root,
	})
	return result
}

func helper(nodeList []*TreeNode) {
	if nodeList == nil || len(nodeList) == 0 {
		return
	}
	temp := make([]int, 0)
	nextArray := make([]*TreeNode, 0)
	for _, node := range nodeList {
		temp = append(temp, node.Val)
		if node.Left != nil {
			nextArray = append(nextArray, node.Left)
		}
		if node.Right != nil {
			nextArray = append(nextArray, node.Right)
		}
	}
	if len(temp) > 0 {
		result = append([][]int{temp}, result...)
	}
	helper(nextArray)
}
