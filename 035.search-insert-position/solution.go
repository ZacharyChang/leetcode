package leetcode

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	start, end := 0, len(nums)-1
	if target <= nums[start] {
		return 0
	}
	if target > nums[end] {
		return end + 1
	}
	for start < end-1 {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			start = mid
		} else {
			end = mid
		}
	}
	return end
}
