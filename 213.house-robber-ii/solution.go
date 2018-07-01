package leetcode

func rob(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	nums1 := make([]int, len(nums)-1)
	copy(nums1, nums)
	nums2 := make([]int, len(nums)-1)
	copy(nums2, nums[1:])
	for i := 1; i < len(nums)-1; i++ {
		if i == 1 {
			nums1[i] = max(nums1[i], nums1[i-1])
			nums2[i] = max(nums2[i], nums2[i-1])
		} else {
			nums1[i] = max(nums1[i-2]+nums1[i], nums1[i-1])
			nums2[i] = max(nums2[i-2]+nums2[i], nums2[i-1])
		}
	}
	return max(nums1[len(nums1)-1], nums2[len(nums2)-1])
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
