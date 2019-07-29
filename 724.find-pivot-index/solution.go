package leetcode

func pivotIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	cur := 0
	for i, v := range nums {
		cur += v
		if cur-v == sum-cur {
			return i
		}
	}
	return -1
}
