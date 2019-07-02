package leetcode

func findLengthOfLCIS(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	res := 1
	tmp := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			tmp++
			if tmp > res {
				res = tmp
			}
		} else {
			tmp = 1
		}
	}
	return res
}
