package leetcode

import (
	"strconv"
)

func addBinary(a string, b string) string {
	carry := 0
	result := ""
	i := len(a) - 1
	j := len(b) - 1
	for ; i >= 0 && j >= 0; i, j = i-1, j-1 {
		sum := byte2Int(a[i]) + byte2Int(b[j]) + carry
		carry = sum / 2
		result = strconv.Itoa(sum%2) + result
	}
	if i >= 0 {
		return addBinary(a[:i+1], strconv.Itoa(carry)) + result
	} else if j >= 0 {
		return addBinary(b[:j+1], strconv.Itoa(carry)) + result
	}
	if carry != 0 {
		return strconv.Itoa(carry) + result
	}
	return result
}

func byte2Int(a byte) int {
	if a == '1' {
		return 1
	}
	return 0
}
