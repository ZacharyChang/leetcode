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

var nodeArray []*TreeNode

func rightSideView(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	nodeArray = append(nodeArray, root)
	for len(nodeArray) > 0 {
		newArray := make([]*TreeNode, 0)
		for _, v := range nodeArray {
			if v.Left != nil {
				newArray = append(newArray, v.Left)
			}
			if v.Right != nil {
				newArray = append(newArray, v.Right)
			}
		}
		result = append(result, nodeArray[len(nodeArray)-1].Val)
		nodeArray = newArray
	}
	return result
}
