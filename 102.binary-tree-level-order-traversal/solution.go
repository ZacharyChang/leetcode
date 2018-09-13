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

func levelOrder(root *TreeNode) [][]int {
	result = make([][]int, 0)
	helper([]*TreeNode{
		root,
	})
	return result
}

func helper(nodeList []*TreeNode) {
	if len(nodeList) == 0 {
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
	result = append(result, temp)
	helper(nextArray)
}
