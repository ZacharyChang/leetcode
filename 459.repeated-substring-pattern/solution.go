package leetcode

func repeatedSubstringPattern(s string) bool {
	if len(s) <= 1 {
		return false
	}
	// 优化
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == s[0] {
			if checkSubstring(s, s[i:]) {
				return true
			}
		}
	}
	return false
}

// bf
func repeatedSubstringPattern_bf(s string) bool {
	if len(s) <= 1 {
		return false
	}
	// 跳过检查自身
	for i := 0; i < len(s)-1; i++ {
		if checkSubstring(s, s[:i+1]) {
			return true
		}
	}
	return false
}

func checkSubstring(s string, sub string) bool {
	if len(s)%len(sub) != 0 {
		return false
	}
	for i := 0; i+len(sub) < len(s)+1; i += len(sub) {
		if s[i:i+len(sub)] != sub {
			return false
		}
	}
	return true
}
