package leetcode

func longestPalindrome(s string) int {
	result := 0
	palindromeMap := make(map[rune]int, 0)
	for _, v := range s {
		palindromeMap[v]++
	}
	for _, v := range palindromeMap {
		// 回文中间可以为单个字符
		if v%2 == 1 && result%2 == 0 {
			result++
		}
		result += v / 2 * 2
	}
	return result
}
