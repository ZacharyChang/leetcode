package leetcode

// Dynamic Programming 动态规划
func wiggleMaxLength_dp(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	dpInc := make([]int, len(nums))
	dpDes := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if dpInc[i] == 0 {
			dpInc[i] = 1
		}
		if dpDes[i] == 0 {
			dpDes[i] = 1
		}
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[i] {
				// 注意DP规律
				// 这里的dpInc[j]表示到j位置时首差值为正的摆动子序列的最大长度
				// 而dpDes[i]为到i位置之前的首差值为负的摆动子序列最大长度
				// 故到j位置的首差值为正的摆动子序列一定有一个长度为dpDes[i]+1
				// nums[...i...j...]，nums[j]>nums[i]
				dpInc[j] = max(dpInc[j], dpDes[i]+1)
			} else if nums[j] < nums[i] {
				dpDes[j] = max(dpDes[j], dpInc[i]+1)
			}
		}
	}
	return max(dpInc[len(dpInc)-1], dpDes[len(dpDes)-1])
}
