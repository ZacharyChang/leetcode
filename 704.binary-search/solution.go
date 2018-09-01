package leetcode

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	start, end := 0, len(nums)-1
	if target < nums[start] || target > nums[end] {
		return -1
	}
	if nums[start] == target {
		return start
	}
	if nums[end] == target {
		return end
	}
	for start < end-1 {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			end = mid
		} else {
			start = mid
		}
	}
	return -1
}
