package leetcode

import "strings"

func uncommonFromSentences(A string, B string) []string {
	res := make([]string, 0)
	senMap := make(map[string]int, 0)
	arr := strings.Split(A+" "+B, " ")
	for _, v := range arr {
		senMap[v]++
	}
	for k := range senMap {
		if senMap[k] == 1 {
			res = append(res, k)
		}
	}
	return res
}
