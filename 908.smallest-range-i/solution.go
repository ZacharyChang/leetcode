package leetcode

import "sort"

func smallestRangeI(A []int, K int) int {
	minA, maxA := min(A), max(A)
	if minA+K >= maxA-K {
		// when the min and max element have common range
		// after min element add K and max element sub K
		return 0
	}
	return (maxA - K) - (minA + K)
}

func min(nums []int) int {
	sort.Ints(nums)
	return nums[0]
}

func max(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)-1]
}
