package leetcode

// Hash Table
func intersection(nums1 []int, nums2 []int) []int {
	result := make([]int, 0)
	resultMap := make(map[int]int, 0)
	map1 := make(map[int]int, 0)
	map2 := make(map[int]int, 0)
	for i := 0; i < len(nums1); i++ {
		map1[nums1[i]]++
	}
	for i := 0; i < len(nums2); i++ {
		map2[nums2[i]]++
	}
	for k := range map1 {
		if map2[k] > 0 {
			if resultMap[k] == 0 {
				result = append(result, k)
				resultMap[k]++
			}
		}
	}
	return result
}
