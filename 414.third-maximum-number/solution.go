package leetcode

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

var max []int

func thirdMax(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	max = []int{INT_MIN, INT_MIN, INT_MIN}
	for i := 0; i < len(nums); i++ {
		push(nums[i])
	}
	if max[0] == INT_MIN {
		return max[2]
	}
	return max[0]
}

func push(n int) {
	if n > max[0] && n < max[1] {
		max[0] = n
	} else if n > max[1] && n < max[2] {
		max[0] = max[1]
		max[1] = n
	} else if n > max[2] {
		max[0] = max[1]
		max[1] = max[2]
		max[2] = n
	}
}
