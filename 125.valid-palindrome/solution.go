package leetcode

import (
	"strings"
)

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	array := []byte(strings.ToLower(s))
	i, j := 0, len(array)-1
	for i < j {
		if !isCharacter(array[i]) {
			i++
		} else if !isCharacter(array[j]) {
			j--
		} else if array[i] == array[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

func isCharacter(b byte) bool {
	if b >= 'A' && b <= 'Z' {
		return true
	}
	if b >= 'a' && b <= 'z' {
		return true
	}
	if b >= '0' && b <= '9' {
		return true
	}
	return false
}
