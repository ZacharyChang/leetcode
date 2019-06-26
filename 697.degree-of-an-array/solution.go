package leetcode

import (
	"math"
)

func findShortestSubArray(nums []int) int {
	numMap := make(map[int][]int, 0)
	res := math.MaxInt64
	degree := 0
	for i, v := range nums {
		numMap[v] = append(numMap[v], i)
		if len(numMap[v]) > degree {
			degree = len(numMap[v])
		}
	}
	for _, v := range numMap {
		if len(v) == degree {
			res = min(res, v[len(v)-1]-v[0]+1)
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
