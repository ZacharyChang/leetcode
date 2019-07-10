package leetcode

func transpose(A [][]int) [][]int {
	res := make([][]int, len(A[0]))
	for i := 0; i < len(A[0]); i++ {
		res[i] = make([]int, len(A))
		for j := 0; j < len(A); j++ {
			res[i][j] = A[j][i]
		}
	}
	return res
}
