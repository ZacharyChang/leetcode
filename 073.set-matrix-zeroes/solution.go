package leetcode

func setZeroes(matrix [][]int) {
	if matrix == nil || len(matrix) == 0 {
		return
	}
	row := make([]int, len(matrix))
	col := make([]int, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				row[i] = 1
				col[j] = 1
			}
		}
	}
	for i := 0; i < len(row); i++ {
		if row[i] == 1 {
			for j := 0; j < len(matrix[0]); j++ {
				matrix[i][j] = 0
			}
		}
	}
	for i := 0; i < len(col); i++ {
		if col[i] == 1 {
			for j := 0; j < len(matrix); j++ {
				matrix[j][i] = 0
			}
		}
	}
}
