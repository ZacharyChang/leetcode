package leetcode

func minMoves(nums []int) int {
	min := nums[0]
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		sum += nums[i]
		if nums[i] < min {
			min = nums[i]
		}
	}
	return sum - min*len(nums)
}
