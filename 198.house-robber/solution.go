package leetcode

func rob(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	for i := 1; i < len(nums); i++ {
		if i == 1 {
			nums[i] = max(nums[i], nums[i-1])
		} else {
			nums[i] = max(nums[i-2]+nums[i], nums[i-1])
		}
	}
	return nums[len(nums)-1]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
