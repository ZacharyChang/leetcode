package leetcode

import "sort"

// Binary Search
func intersection_3(nums1 []int, nums2 []int) []int {
	result := make([]int, 0)
	countMap := make(map[int]int, 0)
	sort.Ints(nums2)
	for i := 0; i < len(nums1); i++ {
		if binarySearch(nums2, nums1[i]) {
			countMap[nums1[i]]++
		}
	}
	for k := range countMap {
		result = append(result, k)
	}
	return result
}

func binarySearch(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	// nums must be sorted
	start, end := 0, len(nums)-1
	if nums[start] == target || nums[end] == target {
		return true
	}
	for start < end-1 {
		mid := (start + end) / 2
		if nums[mid] == target {
			return true
		} else if nums[mid] > target {
			end = mid
		} else {
			start = mid
		}
	}
	return false
}
