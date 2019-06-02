package leetcode

import (
	"sort"
)

// binary search
// time complexity: O(n^2logn)
func triangleNumber(nums []int) int {
	count := 0
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			// binary search to get count from nums[j] to nums[i]+nums[j]
			// because c < a + b
			pos := binarySearch(nums, j+1, len(nums)-1, nums[i]+nums[j])
			count += pos - (j + 1)
		}
	}
	return count
}

func binarySearch(nums []int, left int, right int, target int) int {
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
