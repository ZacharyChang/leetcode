package leetcode

func minSteps(n int) int {
	if n <= 1 {
		return 0
	}
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = i
	}
	for i := 1; i <= n; i++ {
		for j := i * 2; j <= n; j += i {
			dp[j] = min(dp[j], dp[i]+j/i)
		}
	}
	return dp[n]
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
