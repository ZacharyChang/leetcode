/*

Write a function to find the longest common prefix string amongst an array of strings.

*/
package leetcode

import (
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 1; i <= len(strs[0]); i++ {
		for _, str := range strs {
			if strings.HasPrefix(str, strs[0][:i]) {
				continue
			} else {
				if i == 1 {
					return ""
				}
				return strs[0][:i-1]
			}
		}
		// Deal with the situation when the first string is the common prefix
		if i == len(strs[0]) {
			return strs[0]
		}
	}
	return ""
}
