package leetcode

func moveZeroes(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	l := len(nums)
	for i := 0; i < l; i++ {

		if nums[i] == 0 {
			copy(nums[i:len(nums)-1], nums[i+1:])
			nums[len(nums)-1] = 0
			i--
			l--
		}
	}
	return nums
}
