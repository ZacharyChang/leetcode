package leetcode

const (
	intMax = int(^uint(0) >> 1)
	intMin = ^intMax
)

func maxRotateFunction(A []int) int {
	if len(A) <= 1 {
		return 0
	}
	max := intMin
	sum := make([]int, len(A))
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A); j++ {
			sum[i] += j * A[(i+j-1+len(A))%len(A)]
		}
		if sum[i] > max {
			max = sum[i]
		}
	}
	return max
}
