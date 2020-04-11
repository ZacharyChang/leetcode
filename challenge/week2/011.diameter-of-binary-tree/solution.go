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

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// depthMap stores depth of each TreeNode
	depthMap := make(map[*TreeNode]int)
	return helper(root, depthMap)
}

func helper(root *TreeNode, depthMap map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	res := depth(root.Left, depthMap) + depth(root.Right, depthMap)
	res = max(res, helper(root.Left, depthMap))
	res = max(res, helper(root.Right, depthMap))
	return res
}

func depth(root *TreeNode, depthMap map[*TreeNode]int) int {
	if val, exist := depthMap[root]; exist {
		return val
	}
	res := 0
	if root == nil {
		res = 0
	} else if root.Left == nil && root.Right == nil {
		res = 1
	} else {
		res = max(depth(root.Left, depthMap), depth(root.Right, depthMap)) + 1
	}
	depthMap[root] = res
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
