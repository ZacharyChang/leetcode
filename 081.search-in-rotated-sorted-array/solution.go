package leetcode

// Binary Search
func search(nums []int, target int) bool {
	if nums == nil || len(nums) == 0 {
		return false
	}
	start, end := 0, len(nums)-1
	if nums[start] == target || nums[end] == target {
		return true
	}
	for start < end-1 {
		mid := (start + end) / 2
		if nums[mid] == target || nums[start] == target {
			return true
		}
		if nums[mid] == nums[start] {
			start++
			continue
		} else if nums[mid] < nums[start] { // pivot在nums[:mid](左半边)中，右半边为顺序递增
			if between(target, nums[mid], nums[end]) {
				start = mid
			} else {
				end = mid
			}
		} else if nums[mid] > nums[start] { // pivot在nums[mid:](右半边)中，左半边为顺序递增
			if between(target, nums[start], nums[mid]) {
				end = mid
			} else {
				start = mid
			}
		}

	}
	return false
}

func between(a int, b int, c int) bool {
	if a > b && a < c {
		return true
	}
	return false
}
