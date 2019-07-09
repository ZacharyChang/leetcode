package leetcode

import (
	"strings"
)

func rotateString(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	if A == B {
		return true
	}
	for i := 1; i <= len(A); i++ {
		if strings.Index(B, A[:i]) >= 0 {
			if A[i:] == B[:len(B)-i] {
				return true
			}
		}
	}
	return false
}
