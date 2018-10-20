package leetcode

import (
	"bytes"
	"strconv"
	"strings"
)

func decodeString(s string) string {
	if !strings.Contains(s, "[") {
		return s
	}
	var result bytes.Buffer
	tmp := 0
	begin := 0
	numLen := 0
	strLen := 0
	for index, v := range s {
		if v == '[' {
			tmp++
		} else if v == ']' {
			tmp--
			if tmp == 0 {
				count, _ := strconv.Atoi(s[begin : begin+numLen])
				for i := 0; i < count; i++ {
					result.WriteString(decodeString(s[begin+numLen+1 : index]))
				}
				begin = index + 1
				strLen = 0
				numLen = 0
			}
		} else {
			if tmp == 0 {
				if v <= 'z' && v >= 'a' {
					result.WriteRune(v)
					begin++
				} else {
					numLen++
				}
			} else {
				strLen++
			}
		}
	}
	return result.String()
}
