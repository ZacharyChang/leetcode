package leetcode

// Hash Table
// TODO: Stack
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))
	numMap := make(map[int]int, 0)
LOOP:
	for i, v := range nums2 {
		if i == len(nums2)-1 {
			numMap[v] = -1
			continue
		}
		for j := i + 1; j < len(nums2); j++ {
			if nums2[j] > nums2[i] {
				numMap[v] = nums2[j]
				continue LOOP
			}
		}
		numMap[v] = -1
	}
	for i, v := range nums1 {
		res[i] = numMap[v]
	}
	return res
}
