package leetcode

func minSubArrayLen(s int, nums []int) int {
	res := len(nums)
	start := 0
	end := 0
	sum := 0
	for end < len(nums) {
		for sum < s && end < len(nums) {
			sum += nums[end]
			end++
		}
		res = min(res, end-start)
		for sum >= s && start < len(nums) {
			sum -= nums[start]
			start++
			if sum >= s {
				res = min(res, end-start)
			}
		}
	}
	if end-start == len(nums) && sum < s {
		return 0
	}
	return res
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
