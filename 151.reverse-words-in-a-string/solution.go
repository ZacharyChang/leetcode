package leetcode

import (
	"strings"
)

func reverseWords(s string) string {
	arr := strings.Split(s, " ")
	res := make([]string, 0)
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] != "" {
			res = append(res, arr[i])
		}
	}
	return strings.Join(res, " ")
}
