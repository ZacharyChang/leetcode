package leetcode

func maxProduct(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	localMin, localMax, globalMax := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		localMaxCopy := localMax
		localMax = max(max(localMin*nums[i], localMax*nums[i]), nums[i])
		localMin = min(min(localMin*nums[i], localMaxCopy*nums[i]), nums[i])
		globalMax = max(globalMax, localMax)
	}
	return globalMax
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
