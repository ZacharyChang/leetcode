package leetcode

func findBottomLeftValue_2(root *TreeNode) int {
	tmp := root.Val
	array := []*TreeNode{
		root,
	}
	for len(array) > 0 {
		tmp = array[0].Val
		arrayCopy := make([]*TreeNode, len(array))
		copy(arrayCopy, array)
		array = make([]*TreeNode, 0)
		for i := 0; i < len(arrayCopy); i++ {
			if arrayCopy[i].Left != nil {
				array = append(array, arrayCopy[i].Left)
			}
			if arrayCopy[i].Right != nil {
				array = append(array, arrayCopy[i].Right)
			}
		}
	}
	return tmp
}
