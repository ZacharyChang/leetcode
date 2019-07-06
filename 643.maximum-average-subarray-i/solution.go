package leetcode

func findMaxAverage(nums []int, k int) float64 {
	if k >= len(nums) {
		return float64(sum(nums)) / float64(len(nums))
	}
	s := sum(nums[:k])
	maxSum := s
	for i := k; i < len(nums); i++ {
		s += nums[i]
		s -= nums[i-k]
		if s > maxSum {
			maxSum = s
		}
	}
	return float64(maxSum) / float64(k)
}

func sum(nums []int) int {
	res := 0
	for _, v := range nums {
		res += v
	}
	return res
}
