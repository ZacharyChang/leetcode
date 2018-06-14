package leetcode

// Brute Force
func findDuplicate(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	i, j := 0, 0
	for ; i < len(nums); i++ {
		for j = i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return nums[i]
			}
		}
	}
	return nums[i]
}

// Binary Search
// TODO
func findDuplicate_2(nums []int) int {
	return 0
}
