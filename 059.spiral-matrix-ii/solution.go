package leetcode

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, n)
	}
	start := 1
	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			res[i][j] = start
			start++
		}
		for j := i; j < n-i-1; j++ {
			res[j][n-i-1] = start
			start++
		}
		for j := n - i - 1; j > i; j-- {
			res[n-i-1][j] = start
			start++
		}
		for j := n - i - 1; j > i; j-- {
			res[j][i] = start
			start++
		}
	}
	if n%2 == 1 {
		res[n/2][n/2] = start
	}
	return res
}
