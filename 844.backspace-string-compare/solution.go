package leetcode

import (
	"bytes"
)

// build string and compare
func backspaceCompare(S string, T string) bool {
	return str2Res(S) == str2Res(T)
}

func str2Res(S string) string {
	var res bytes.Buffer
	for _, v := range S {
		if v == '#' {
			if res.Len() == 0 {
				continue
			} else {
				res.Truncate(res.Len() - 1)
			}
		} else {
			res.WriteRune(v)
		}
	}
	return res.String()
}
