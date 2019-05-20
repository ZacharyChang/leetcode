package leetcode

// dp solution
// dp[i][i] = true
// dp[i][i+1] = s[i] == s[i+1]
// dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
// time complexity: O(n^2)
func longestPalindrome_2(s string) string {
	if len(s) <= 1 {
		return s
	}
	dp := make([][]bool, len(s))
	res := s[0:1]
	// init dp array
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
		if i < len(s)-1 && s[i] == s[i+1] {
			dp[i][i+1] = true
			res = s[i : i+2]
		}
	}
	// dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
	for i := len(s) - 1; i >= 0; i-- {
		// i+1 <= j-1
		for j := i + 2; j < len(s); j++ {
			dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
			if dp[i][j] && j-i+1 > len(res) {
				res = s[i : j+1]
			}
		}
	}
	return res
}
