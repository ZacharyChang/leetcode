package leetcode

// DP
func maxSubArray(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	max := nums[0]
	dpSum := nums[0]
	// dpArray := make([]int, len(nums))
	// dpArray[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dpSum > 0 {
			dpSum = dpSum + nums[i]
		} else {
			dpSum = nums[i]
		}
		if dpSum > max {
			max = dpSum
		}
		// if dpArray[i-1] > 0 {
		// 	dpArray[i] = dpArray[i-1] + nums[i]
		// } else {
		// 	dpArray[i] = nums[i]
		// }
		// if dpArray[i] > max {
		// 	max = dpArray[i]
		// }
	}
	return max
}

// Divide and Conquer
func maxSubArray_2(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	// TODO
	return nums[0]
}
