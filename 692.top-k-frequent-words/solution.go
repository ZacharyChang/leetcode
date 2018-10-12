package leetcode

import (
	"sort"
)

func topKFrequent(words []string, k int) []string {
	result := make([]string, 0)
	array := make([][]string, len(words)+1)
	wordMap := make(map[string]int, len(words))
	for _, v := range words {
		wordMap[v]++
	}
	for k, v := range wordMap {
		array[v] = append(array[v], k)
	}
	for i := len(array) - 1; i >= 0; i-- {
		if len(array[i]) > 0 {
			sort.Strings(array[i])
			result = append(result, array[i]...)
			if len(result) >= k {
				return result[:k]
			}
		}
	}
	return result
}
