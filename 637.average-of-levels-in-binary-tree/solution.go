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

func averageOfLevels(root *TreeNode) []float64 {
	result := make([]float64, 0)
	nodeArray := []*TreeNode{
		root,
	}
	for len(nodeArray) > 0 {
		sum := 0
		newArray := make([]*TreeNode, 0)
		for _, v := range nodeArray {
			sum += v.Val
			if v.Left != nil {
				newArray = append(newArray, v.Left)
			}
			if v.Right != nil {
				newArray = append(newArray, v.Right)
			}
		}
		result = append(result, float64(sum)/float64(len(nodeArray)))
		nodeArray = newArray
	}
	return result
}
