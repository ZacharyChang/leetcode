package leetcode

import (
	"sort"
)

func hIndex(citations []int) int {
	sort.Ints(citations)
	res := 0
	for i := 0; i < len(citations); i++ {
		curH := min(citations[i], len(citations)-i)
		res = max(res, curH)
	}
	return res
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
