package leetcode

import (
	"bytes"
	"strings"
)

func frequencySort(s string) string {
	freqMap := make(map[rune]int, len(s))
	var result bytes.Buffer
	for _, v := range s {
		freqMap[v]++
	}
	array := make([]string, len(s)+1)
	for k, v := range freqMap {
		array[v] += strings.Repeat(string(k), v)
	}
	for i := len(array) - 1; i >= 0; i-- {
		if array[i] != "" {
			result.WriteString(array[i])
		}
	}
	return result.String()
}
