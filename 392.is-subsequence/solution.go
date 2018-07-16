package leetcode

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	i, j := 0, 0
	for ; i < len(s) && j < len(t); j++ {
		if s[i] == t[j] {
			i++
		}
	}
	if i == len(s) {
		return true
	}
	return false
}
