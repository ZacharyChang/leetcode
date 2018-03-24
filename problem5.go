// Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

// Example:

// Input: "babad"

// Output: "bab"

// Note: "aba" is also a valid answer.

// Example:

// Input: "cbbd"

// Output: "bb"

package leetcode

func longestPalindrome(s string) string {
	max := 0
	maxStr := ""
	// ship the situation where length will be less than 'max'
	for i := 0; i < len(s)-max; i++ {
		// just start from the length more than value 'max'
		for j := i + max + 1; j <= len(s); j++ {
			// length is j-i
			subStr := s[i:j]
			if isPalindrome(subStr) {
				max = len(subStr)
				maxStr = subStr
			}
		}
	}
	return maxStr
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
