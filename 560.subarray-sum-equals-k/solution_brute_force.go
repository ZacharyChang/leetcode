package leetcode

// brute force
func subarraySum1(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	sum := make([]int, len(nums)+1)
	count := 0
	sum[0] = 0
	for i := 0; i < len(nums); i++ {
		sum[i+1] = sum[i] + nums[i]
	}
	for i := 0; i <= len(nums); i++ {
		for j := i + 1; j <= len(nums); j++ {
			if sum[j]-sum[i] == k {
				count++
			}
		}
	}
	return count
}
