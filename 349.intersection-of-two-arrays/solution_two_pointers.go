package leetcode

import "sort"

// Two Pointers
func intersection_2(nums1 []int, nums2 []int) []int {
	result := make([]int, 0)
	sort.Ints(nums1)
	sort.Ints(nums2)
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			if len(result) == 0 || result[len(result)-1] != nums1[i] {
				result = append(result, nums1[i])
			}
			i++
			j++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			i++
		}
	}
	return result
}
