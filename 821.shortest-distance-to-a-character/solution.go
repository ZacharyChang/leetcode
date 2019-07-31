package leetcode

import "math"

func shortestToChar(S string, C byte) []int {
	res := make([]int, len(S))
	for i := range res {
		res[i] = math.MaxInt32
	}
	for i := range S {
		if S[i] == C {
			res[i] = 0
			for j := i - 1; j >= 0; j-- {
				res[j] = min(res[j], i-j)
			}
			for j := i+1; j <len(S); j++ {
				res[j] = min(res[j], j-i)
			}
		}
	}
	return res
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
