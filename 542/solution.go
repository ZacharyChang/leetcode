package leetcode

import (
	"math"
)

func updateMatrix(matrix [][]int) [][]int {
	// init the values of '1' to MAX
	// use MaxInt32 here to prevent int overflow
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 1 {
				matrix[i][j] = math.MaxInt32
			}
		}
	}

	for m := 0; m < len(matrix) || m < len(matrix[0]); m++ {
		count := 0
		for i := 0; i < len(matrix); i++ {
			for j := 0; j < len(matrix[0]); j++ {
				// math.MaxInt64+1 will cause overflow
				if i-1 >= 0 && matrix[i-1][j] > matrix[i][j]+1 {
					matrix[i-1][j] = matrix[i][j] + 1
					count++
				}
				if i+1 < len(matrix) && matrix[i+1][j] > matrix[i][j]+1 {
					matrix[i+1][j] = matrix[i][j] + 1
					count++
				}
				if j-1 >= 0 && matrix[i][j-1] > matrix[i][j]+1 {
					matrix[i][j-1] = matrix[i][j] + 1
					count++
				}
				if j+1 < len(matrix[0]) && matrix[i][j+1] > matrix[i][j]+1 {
					matrix[i][j+1] = matrix[i][j] + 1
					count++
				}
			}
		}
		if count == 0 {
			return matrix
		}
	}
	return matrix
}

