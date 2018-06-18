package leetcode

// Binary Search
func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	start := 0
	end := len(nums) - 1
	for start < end-1 {
		mid := (start + end) / 2
		if nums[mid] < nums[mid-1] {
			end = mid
		} else if nums[mid] < nums[mid+1] {
			start = mid
		} else {
			return mid
		}
	}
	if nums[start] < nums[end] {
		return end
	}
	return start
}
