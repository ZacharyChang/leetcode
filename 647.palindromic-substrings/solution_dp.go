package leetcode

// dp
// time complexity: O(n^2)
// dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
func countSubstrings_2(s string) int {
	res := 0
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
		res++
	}
	for i := len(s) - 2; i >= 0; i-- {
		dp[i][i+1] = s[i] == s[i+1]
		if dp[i][i+1] {
			res++
		}
		// i+1<=j-1
		for j := i + 2; j < len(s); j++ {
			dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
			if dp[i][j] {
				res++
			}
		}
	}
	return res
}
