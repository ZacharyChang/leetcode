package leetcode

func numberOfArithmeticSlices_bf(A []int) int {
	count := 0
	for i := 0; i < len(A); i++ {
		for j := i + 2; j < len(A); j++ {
			if isArithmetic(A[i : j+1]) {
				count++
			}
		}
	}
	return count
}

func isArithmetic(nums []int) bool {
	diff := (nums[len(nums)-1] - nums[0]) / (len(nums) - 1)
	for i := 0; i < len(nums); i++ {
		if nums[i] == nums[0]+diff*i {
			continue
		}
		return false
	}
	return true
}
