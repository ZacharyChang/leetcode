package leetcode

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	result := 0
	// dp[i][j]为以matrix[i][j]为右下角顶点的最大正方形边长
	dp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0]))
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else if i > 0 && j > 0 {
					// dp递推
					dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
				}
			}
			if dp[i][j] > result {
				result = dp[i][j]
			}
		}
	}
	return result * result
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
