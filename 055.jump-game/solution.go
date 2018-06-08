package leetcode

// 贪心
func canJump(nums []int) bool {
	if nums == nil || len(nums) == 0 {
		return true
	}
	if nums[0] >= len(nums)-1 {
		return true
	}
	max := nums[0]
	for i := 0; i < len(nums)-1; i++ {
		// 检查并更新最远距离，注意条件 max>=i 表示i为可访问
		if max >= i && i+nums[i] > max {
			max = i + nums[i]
		}
		if max >= len(nums)-1 {
			return true
		}
	}
	return false
}

// 递归
func canJump_2(nums []int) bool {
	if nums == nil || len(nums) == 0 {
		return true
	}
	if nums[0] >= len(nums)-1 {
		return true
	}
	for i := nums[0]; i >= 1; i-- {
		if canJump(nums[i:]) {
			return true
		}
	}
	return false
}
