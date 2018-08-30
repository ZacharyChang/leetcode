package leetcode

import "sort"

// Binary Search
func intersect_2(nums1 []int, nums2 []int) []int {
	result := make([]int, 0)
	countMap := make(map[int]int, 0)
	sort.Ints(nums2)
	for i := 0; i < len(nums1); i++ {
		if pos := binarySearch(nums2, nums1[i]); pos >= 0 {
			countMap[nums1[i]]++
			nums2 = append(nums2[:pos], nums2[pos+1:]...)
		}
	}
	for k, v := range countMap {
		for i := 0; i < v; i++ {
			result = append(result, k)
		}
	}
	return result
}

func binarySearch(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	// nums must be sorted
	start, end := 0, len(nums)-1
	if nums[start] == target {
		return start
	}
	if nums[end] == target {
		return end
	}
	for start < end-1 {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			end = mid
		} else {
			start = mid
		}
	}
	return -1
}
