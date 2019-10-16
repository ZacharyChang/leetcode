package leetcode

func numTrees(n int) int {
	if n <= 1 {
		return 1
	}
	nums := make([]int, n+1)
	nums[0] = 1
	nums[1] = 1
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			nums[i] += nums[j] * nums[i-1-j]
		}
	}
	return nums[n]
}
