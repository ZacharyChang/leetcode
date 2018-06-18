package leetcode

func productExceptSelf(nums []int) []int {
	if nums == nil || len(nums) == 0 {
		return []int{}
	}
	result := make([]int, len(nums))
	leftProductArray := make([]int, len(nums))
	leftProductArray[0] = 1
	rightProductArray := make([]int, len(nums))
	rightProductArray[len(nums)-1] = 1
	for i := 1; i < len(nums); i++ {
		leftProductArray[i] = leftProductArray[i-1] * nums[i-1]
		rightProductArray[len(nums)-1-i] = rightProductArray[len(nums)-i] * nums[len(nums)-i]
	}
	for j := 0; j < len(nums); j++ {
		result[j] = leftProductArray[j] * rightProductArray[j]
	}
	return result
}
