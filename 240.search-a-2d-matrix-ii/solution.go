package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	if len(matrix[0]) == 0 {
		return false
	}
	i, j := 0, len(matrix)-1
	for i < j-1 {
		mid := (i + j) / 2
		if matrix[mid][0] == target {
			return true
		} else if matrix[mid][0] < target {
			i = mid
		} else {
			j = mid
		}
	}
	m, n := 0, len(matrix[0])-1
	for m < n-1 {
		mid := (m + n) / 2
		if matrix[0][mid] == target {
			return true
		} else if matrix[0][mid] < target {
			m = mid
		} else {
			n = mid
		}
	}
	return search(matrix, j, n, target)
}

func search(matrix [][]int, m int, n int, target int) bool {
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if matrix[i][j] == target {
				return true
			}
		}
	}
	return false
}
