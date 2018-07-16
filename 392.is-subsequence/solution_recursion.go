package leetcode

import "strings"

func isSubsequence_bf(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(t) == 0 {
		return false
	}
	if s[0] == t[0] {
		return isSubsequence_bf(s[1:], t[1:])
	} else {
		return isSubsequence_bf(s, t[1:])
	}
}

func isSubsequence_strings(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(t) == 0 {
		return false
	}
	index := strings.Index(t, s[:1])
	if index >= 0 {
		return isSubsequence_strings(s[1:], t[index+1:])
	}
	return false
}
