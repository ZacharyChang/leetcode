package leetcode

func countNumbersWithUniqueDigits_dp(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 10
	}
	dp := make([]int, n+1)
	result := make([]int, n+1)
	dp[0] = 1
	dp[1] = 9
	result[0] = dp[0]
	result[1] = result[0] + dp[1]
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] * (10 - (i - 1))
		result[i] = result[i-1] + dp[i]
	}
	return result[n]
}
