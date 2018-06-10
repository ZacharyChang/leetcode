package leetcode

// Binary Search
func search(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	if nums[0] == target {
		return 0
	}
	if nums[len(nums)-1] == target {
		return len(nums) - 1
	}
	start := 0
	end := len(nums) - 1
	for start < end-1 {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		}
		// pivot在nums[mid:](右半边)中，左半边为顺序递增
		if nums[mid] > nums[end] {
			if target > nums[start] && target < nums[mid] {
				end = mid
			} else {
				start = mid
			}
		} else if nums[mid] < nums[start] { // pivot在nums[:mid](左半边)中，右半边为顺序递增
			if target > nums[mid] && target < nums[end] {
				start = mid
			} else {
				end = mid
			}
		} else { // pivot不在nums[start:end]中
			if nums[mid] > target {
				end = mid
			} else {
				start = mid
			}
		}
	}
	return -1
}
