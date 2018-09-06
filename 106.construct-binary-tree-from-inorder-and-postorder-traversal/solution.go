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

var treeMap map[int]int

func buildTree(inorder []int, postorder []int) *TreeNode {
	treeMap = make(map[int]int, len(inorder))
	for i := 0; i < len(inorder); i++ {
		treeMap[inorder[i]] = i
	}
	return helper(inorder, postorder, 0, len(inorder)-1, 0, len(postorder)-1)
}

func helper(inorder []int, postorder []int, inLeft int, inRight int, postLeft int, postRight int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 || inLeft > inRight || postLeft > postRight {
		return nil
	}
	// 根据postorder获取当前父节点
	rootVal := postorder[postRight]
	// 从map中获取该节点在中序遍历中的位置
	rootPos := treeMap[rootVal]
	// 从inorder中推算左子树节点个数
	leftNum := rootPos - inLeft
	return &TreeNode{
		rootVal,
		helper(inorder, postorder, inLeft, rootPos-1, postLeft, postLeft+leftNum-1),
		helper(inorder, postorder, rootPos+1, inRight, postLeft+leftNum, postRight-1),
	}
}
