// Given a 32-bit signed integer, reverse digits of an integer.

// Example 1:

// Input: 123
// Output:  321
// Example 2:

// Input: -123
// Output: -321
// Example 3:

// Input: 120
// Output: 21
// Note:
// Assume we are dealing with an environment which could only hold integers within the 32-bit signed integer range. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.

package leetcode

import (
	"math"
	"strconv"
)

func reverse(num int) int {
	flag := false
	if num < 0 {
		flag = true
		num = 0 - num
	}
	s := strconv.Itoa(num)
	reversedStr := ""
	for i := 0; i < len(s); i++ {
		reversedStr += string(s[len(s)-1-i])
	}
	reversedNum, _ := strconv.Atoi(reversedStr)
	if float64(num) >= math.Pow(2, 31) || float64(reversedNum) >= math.Pow(2, 31) {
		return 0
	}
	if flag {
		return 0 - reversedNum
	}
	return reversedNum
}
