package leetcode

// Greedy 贪心算法
func wiggleMaxLength(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	// 分为两种情况讨论，取最大值。时间复杂度 O(n)
	// 1. 起始差值为正值
	// 2. 起始差值为负值
	// pos=>positive, nega=>negative
	posCount, negaCount := 1, 1
	posFlag, negaFlag := true, false
	diffArray := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		diffArray[i] = nums[i] - nums[i-1]
		if diffArray[i] > 0 {
			if posFlag {
				posFlag = false
				posCount++
			}
			if negaFlag {
				negaFlag = false
				negaCount++
			}
		} else if diffArray[i] < 0 {
			if !posFlag {
				posFlag = true
				posCount++
			}
			if !negaFlag {
				negaFlag = true
				negaCount++
			}
		}
	}
	return max(posCount, negaCount)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
