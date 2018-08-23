package leetcode

func findMin(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	start := 0
	end := len(nums) - 1
	result := nums[0]
	for start < end-1 {
		mid := (start + end) / 2
		if nums[mid] < nums[end] { // pivot在mid左侧
			end = mid
		} else if nums[mid] > nums[end] { // pivot在mid右侧
			start = mid
		} else { // 不确定pivot位置
			end--
		}
	}
	result = min(result, nums[end])
	return result
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
