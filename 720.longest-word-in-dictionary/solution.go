package leetcode

import (
	"sort"
)

func longestWord(words []string) string {

	wordMap := make(map[string]bool, 0)
	sort.Strings(words)
	for _, v := range words {
		wordMap[v] = true
	}
	maxLen := 0
	result := ""
	for _, word := range words {
		if len(word) > maxLen && checkValid(word, &wordMap) {
			result = word
			maxLen = len(word)
		}
	}
	return result
}

func checkValid(s string, mPtr *map[string]bool) bool {
	m := *mPtr
	for i := 0; i < len(s); i++ {
		if !m[s[0:i+1]] {
			return false
		}
	}
	return true
}
