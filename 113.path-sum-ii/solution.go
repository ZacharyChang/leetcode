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

var res [][]int

func pathSum(root *TreeNode, sum int) [][]int {
	res = make([][]int, 0)
	helper(root, sum, 0, []int{})
	return res
}

func helper(root *TreeNode, sum int, cur int, array []int) {
	if root == nil {
		return
	}
	newArray := make([]int, len(array)+1)
	copy(newArray, array)
	newArray[len(newArray)-1] = root.Val
	newSum := root.Val + cur
	if newSum == sum && root.Left == nil && root.Right == nil {
		res = append(res, newArray)
	}
	helper(root.Left, sum, newSum, newArray)
	helper(root.Right, sum, newSum, newArray)
}
