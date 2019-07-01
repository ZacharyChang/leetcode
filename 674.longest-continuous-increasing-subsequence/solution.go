package leetcode

func findLengthOfLCIS(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	res := 1
	array := make([]int, len(nums))
	array[0] = 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			array[i] = array[i-1] + 1
		} else {
			array[i] = 1
		}
		if array[i] > res {
			res = array[i]
		}
	}
	return res
}
