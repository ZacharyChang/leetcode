package leetcode

import "sort"

// brute force
// sort first will make the algorithm faster
// time complexity: O(n^3)
func triangleNumber_2(nums []int) int {
	count := 0
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j] > nums[k] {
					count++
				}
			}
		}
	}
	return count
}
