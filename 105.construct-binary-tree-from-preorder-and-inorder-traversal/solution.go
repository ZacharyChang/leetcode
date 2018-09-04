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

func buildTree(preorder []int, inorder []int) *TreeNode {
	treeMap = make(map[int]int, len(inorder))
	for i := 0; i < len(inorder); i++ {
		treeMap[inorder[i]] = i
	}
	return build(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1)
}

func build(preorder []int, inorder []int, leftPre int, rightPre int, leftIn int, rightIn int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 || leftPre > rightPre || leftIn > rightIn {
		return nil
	}
	// 根据preorder获取当前父节点
	rootVal := preorder[leftPre]
	// 从map中获取该节点在中序遍历中的位置
	rootPos := treeMap[rootVal]
	// 从inorder中推算左子树节点个数
	leftNum := rootPos - leftIn
	return &TreeNode{
		rootVal,
		build(preorder, inorder, leftPre+1, leftPre+leftNum, leftIn, rootPos-1),
		build(preorder, inorder, leftPre+leftNum+1, rightPre, rootPos+1, rightIn),
	}
}
