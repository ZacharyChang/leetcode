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

type TreeNodes []*TreeNode

func (nodes *TreeNodes) append(n *TreeNode) {
	if n != nil {
		*nodes = append(*nodes, n)
	}
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	inOrder := true
	stack := make(TreeNodes, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		line := make([]int, 0)
		newStack := make(TreeNodes, 0)
		for i := 0; i < len(stack); i++ {
			newStack.append(stack[i].Left)
			newStack.append(stack[i].Right)
			if inOrder {
				line = append(line, stack[i].Val)
			} else {
				line = append(line, stack[len(stack)-1-i].Val)
			}
		}
		res = append(res, line)
		stack = newStack
		inOrder = !inOrder
	}
	return res
}
