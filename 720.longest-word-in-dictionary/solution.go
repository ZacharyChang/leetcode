package leetcode

import (
	"sort"
)

func longestWord(words []string) string {

	wordMap := make(map[string]bool, 0)
	wordMap[""] = true
	sort.Strings(words)
	maxLen := 0
	result := ""
	for _, word := range words {
		if checkValid(word, &wordMap) {
			wordMap[word] = true
			if len(word) > maxLen {
				result = word
				maxLen = len(word)
			}

		}
	}
	return result
}

func checkValid(s string, mPtr *map[string]bool) bool {
	m := *mPtr
	return m[s[0:len(s)-1]]
}
