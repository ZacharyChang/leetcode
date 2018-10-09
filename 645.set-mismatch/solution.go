package leetcode

func findErrorNums(nums []int) []int {
	// array is quicker than map
	numArray := make([]int, len(nums)+1)
	sum := len(nums) * (len(nums) + 1) / 2
	dupNum := nums[0]
	for _, v := range nums {
		numArray[v]++
		sum -= v
		if numArray[v] == 2 {
			dupNum = v
		}
	}
	return []int{
		dupNum,
		sum + dupNum,
	}
}
