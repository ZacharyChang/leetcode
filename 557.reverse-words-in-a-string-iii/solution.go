package leetcode

import (
	"bytes"
	"strings"
)

func reverseWords(s string) string {
	array := strings.Split(s, " ")
	for i, v := range array {
		array[i] = reverse(v)
	}
	return strings.Join(array, " ")
}

func reverse(s string) string {
	var buffer bytes.Buffer
	for i := len(s) - 1; i >= 0; i-- {
		buffer.WriteByte(s[i])
	}
	return buffer.String()
}
