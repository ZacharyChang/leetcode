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

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return generateRange(1, n)
}

func generateRange(from int, to int) []*TreeNode {
	res := make([]*TreeNode, 0)
	if from > to {
		res = append(res, nil)
		return res
	} else if from == to {
		res = append(res, &TreeNode{
			Val: from,
		})
		return res
	}
	for i := from; i <= to; i++ {
		leftSub := generateRange(from, i-1)
		rightSub := generateRange(i+1, to)
		for _, v1 := range leftSub {
			for _, v2 := range rightSub {
				res = append(res, &TreeNode{
					Val:   i,
					Left:  v1,
					Right: v2,
				})
			}
		}
	}
	return res
}
