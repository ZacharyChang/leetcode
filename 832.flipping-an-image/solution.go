package leetcode

func flipAndInvertImage(A [][]int) [][]int {
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[0])/2; j++ {
			A[i][j], A[i][len(A[0])-1-j] = A[i][len(A[0])-1-j], A[i][j]
		}
		for j := 0; j < len(A[0]); j++ {
			A[i][j] = 1 - A[i][j]
		}
	}
	return A
}
