package leetcode

// expand around center
// time complexity: O(n^2)
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	res := ""
	max := 0
	for i := 0; i < len(s)-1; i++ {
		left, right := findPalindrome(s, i, i)
		if right-left+1 > max {
			max = right - left + 1
			res = s[left : right+1]
		}
		left, right = findPalindrome(s, i, i+1)
		if right-left+1 > max {
			max = right - left + 1
			res = s[left : right+1]
		}
	}
	return res
}

func findPalindrome(s string, left int, right int) (int, int) {
	if s[left] != s[right] {
		return 0, 0
	}
	for left-1 >= 0 && right+1 < len(s) && s[left-1] == s[right+1] {
		left--
		right++
	}
	return left, right
}
