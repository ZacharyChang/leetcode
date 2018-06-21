package leetcode

func rotate(nums []int, k int) {
	if nums == nil || len(nums) == 0 {
		return
	}
	k = k % len(nums)
	temp := make([]int, k)
	copy(temp[:k], nums[len(nums)-k:])
	copy(nums[k:], nums[:len(nums)-k])
	copy(nums[:k], temp)
}
