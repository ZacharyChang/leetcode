package leetcode

func imageSmoother(M [][]int) [][]int {
	res := make([][]int, len(M))
	count := make([][]int, len(M))
	for i := range res {
		res[i] = make([]int, len(M[0]))
		count[i] = make([]int, len(M[0]))
	}
	for i := 0; i < len(M); i++ {
		for j := 0; j < len(M[0]); j++ {
			if i-1 >= 0 {
				if j-1 >= 0 {
					res[i-1][j-1] += M[i][j]
					count[i-1][j-1]++
				}
				res[i-1][j] += M[i][j]
				count[i-1][j]++
				if j+1 < len(M[0]) {
					res[i-1][j+1] += M[i][j]
					count[i-1][j+1]++
				}
			}
			if i+1 < len(M) {
				if j-1 >= 0 {
					res[i+1][j-1] += M[i][j]
					count[i+1][j-1]++
				}
				res[i+1][j] += M[i][j]
				count[i+1][j]++
				if j+1 < len(M[0]) {
					res[i+1][j+1] += M[i][j]
					count[i+1][j+1]++
				}
			}
			if j-1 >= 0 {
				res[i][j-1] += M[i][j]
				count[i][j-1]++
			}
			res[i][j] += M[i][j]
			count[i][j]++
			if j+1 < len(M[0]) {
				res[i][j+1] += M[i][j]
				count[i][j+1]++
			}
		}
	}
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i]); j++ {
			res[i][j] = res[i][j] / count[i][j]
		}
	}
	return res
}
