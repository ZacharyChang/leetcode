package leetcode

func validPalindrome(s string) bool {
	return helper(s, 1)
}

func helper(s string, skip int) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] == s[right] {
			left++
			right--
		} else {
			if skip == 0 {
				return false
			} else {
				skip--
				if s[left] == s[right-1] && s[left+1] == s[right] {
					return helper(s[left:right], skip) || helper(s[left+1:right+1], skip)
				}
				if s[left] == s[right-1] {
					right--
				} else if s[left+1] == s[right] {
					left++
				} else {
					return false
				}
			}

		}
	}
	return s[left] == s[right]
}
