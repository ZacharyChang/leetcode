package leetcode

// expand around center
// time complexity: O(n^2)
func countSubstrings(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		res += checkPalindromic(s, i, i)
		if i < len(s)-1 {
			res += checkPalindromic(s, i, i+1)
		}
	}
	return res
}

func checkPalindromic(s string, left int, right int) int {
	count := 0
	for left >= 0 && right < len(s) {
		if s[left] != s[right] {
			break
		}
		count++
		left--
		right++
	}
	return count
}
