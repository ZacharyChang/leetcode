package leetcode

import "bytes"

func reverseString(s string) string {
	var buffer bytes.Buffer
	for i := len(s) - 1; i >= 0; i-- {
		buffer.WriteByte(s[i])
	}
	return buffer.String()
}
