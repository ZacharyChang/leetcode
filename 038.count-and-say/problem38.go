// The count-and-say sequence is the sequence of integers with the first five terms as following:

// 1.     1
// 2.     11
// 3.     21
// 4.     1211
// 5.     111221
// 1 is read off as "one 1" or 11.
// 11 is read off as "two 1s" or 21.
// 21 is read off as "one 2, then one 1" or 1211.
// Given an integer n, generate the nth term of the count-and-say sequence.

// Note: Each term of the sequence of integers will be represented as a string.

// Example 1:

// Input: 1
// Output: "1"
// Example 2:

// Input: 4
// Output: "1211"
package leetcode

import (
	"strconv"
)

func countAndSay(n int) string {
	count := 0
	result := ""
	if n == 0 {
		return ""
	}
	if n == 1 {
		return "1"
	}
	lastStr := countAndSay(n - 1)
	temp := lastStr[0]
	for i, _ := range lastStr {
		if lastStr[i] == temp {
			count++
		} else {
			if count > 0 {
				result += strconv.Itoa(count)
				result += string(temp)
				// reset count and temp value
				count = 1
				temp = lastStr[i]
			}
		}
	}
	result += strconv.Itoa(count)
	result += string(temp)

	return result
}
