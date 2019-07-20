package leetcode

import (
	"sort"
)

func shortestCompletingWord(licensePlate string, words []string) string {
	// because the one that occurs first should return first, the sort should be stable
	sort.Stable(wordList(words))

	m := str2Map(licensePlate)
	for _, v := range words {
		if canComplete(v, m) {
			return v
		}
	}
	return ""
}

type wordList []string

func (w wordList) Len() int {
	return len(w)
}

func (w wordList) Less(i int, j int) bool {
	return len(w[i]) < len(w[j])
}

func (w wordList) Swap(i int, j int) {
	w[i], w[j] = w[j], w[i]
}

func str2Map(s string) map[rune]int {
	res := make(map[rune]int, 0)
	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			res[v]++
		}
		if v >= 'A' && v <= 'Z' {
			res[v+('a'-'A')]++
		}
	}
	return res
}

func canComplete(word string, m map[rune]int) bool {
	wordMap := str2Map(word)
	for k, v := range m {
		if wordMap[k] < v {
			return false
		}
	}
	return true
}
