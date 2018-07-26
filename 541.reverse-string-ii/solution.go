package leetcode

import (
	"bytes"
)

func reverseStr(s string, k int) string {
	result := ""
	i := 0
	for ; i+k <= len(s); i += 2 * k {
		result += reverse(s[i : i+k])
		result += s[i+k : min(i+2*k, len(s))]
	}
	if i < len(s) {
		result += reverse(s[i:])
	}
	return result
}

func reverse(s string) string {
	var buffer bytes.Buffer
	for i := len(s) - 1; i >= 0; i-- {
		buffer.WriteByte(s[i])
	}
	return buffer.String()
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
