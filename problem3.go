// Given a string, find the length of the longest substring without repeating characters.

// Examples:

// Given "abcabcbb", the answer is "abc", which the length is 3.

// Given "bbbbb", the answer is "b", with the length of 1.

// Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

package leetcode

import (
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	max := 0
	for i := 0; i < len(s); i++ {
		// begin from the (i+max), just skip the length which is less than max
		for j := i + max; j <= len(s); j++ {
			str := s[i:j]
			// break the loop if the substring already repeats
			if isRepeated(str) {
				break
			} else if len(str) > max {
				max = len(str)
			}
		}
	}
	return max
}

func isRepeated(s string) bool {
	for _, sub := range s {
		if strings.Count(s, string(sub)) > 1 {
			return true
		}
	}
	return false
}
