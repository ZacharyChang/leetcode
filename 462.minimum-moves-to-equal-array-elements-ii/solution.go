package leetcode

import (
	"sort"
)

func minMoves2(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	result := 0
	sort.Ints(nums)
	mid := nums[len(nums)/2]
	for i := 0; i < len(nums); i++ {
		result += abs(nums[i] - mid)
	}
	return result
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
