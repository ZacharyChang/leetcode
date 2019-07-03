package leetcode

import (
	"sort"
	"strings"
)

func replaceWords(dict []string, sentence string) string {
	sentenceArray := strings.Split(sentence, " ")
	var result []string
	sort.Strings(dict)
	wordMap := make(map[string]bool, len(dict))
	for _, v := range dict {
		wordMap[v] = true
	}
	for _, v := range sentenceArray {
		result = append(result, replace(wordMap, v))
	}
	return strings.Join(result, " ")
}

func replace(wordMap map[string]bool, sentence string) string {
	for i := 1; i <= len(sentence); i++ {
		if wordMap[sentence[:i]] {
			return sentence[:i]
		}
	}
	return sentence
}
