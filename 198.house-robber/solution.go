package leetcode

func rob(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	for i := 2; i < len(nums); i++ {
		nums[i] = max(nums[i-2]+nums[i], nums[i-1])
	}
	return nums[len(nums)-1]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
