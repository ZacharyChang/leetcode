package leetcode

func findMin(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	start := 0
	end := len(nums) - 1
	for start < end-1 {
		mid := (start + end) / 2
		if nums[mid] > nums[start] && nums[mid] < nums[end] { // nums[start:end+1]为递增顺序(无pivot)
			return nums[start]
		} else if nums[mid] < nums[start] && nums[mid] < nums[end] { // pivot在nums[mid]左侧或恰好为nums[mid]
			end = mid
		} else if nums[mid] > nums[start] && nums[mid] > nums[end] { // pivot在nums[mid]右侧
			start = mid
		}
	}
	if nums[start] < nums[end] {
		return nums[start]
	}
	return nums[end]
}
