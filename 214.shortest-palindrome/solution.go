package leetcode

// expand around center
func shortestPalindrome(s string) string {
	for i := len(s); i > 0; i-- {
		if checkPalindrome(s[0:i]) {
			return reverse(s[i:]) + s
		}
	}
	return reverse(s) + s
}

func checkPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return s[left] == s[right]
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
