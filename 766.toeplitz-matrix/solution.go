package leetcode

func isToeplitzMatrix(matrix [][]int) bool {
	matrixMap := make(map[int]int, 0)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if val, ok := matrixMap[j-i]; ok {
				if matrix[i][j] != val {
					return false
				}
			} else {
				matrixMap[j-i] = matrix[i][j]
			}
		}
	}
	return true
}
