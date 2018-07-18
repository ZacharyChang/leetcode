package leetcode

func numberOfArithmeticSlices(A []int) int {
	if len(A) < 3 {
		return 0
	}
	count := 0
	array := make([]int, len(A))
	for i := 1; i < len(A); i++ {
		array[i] = A[i] - A[i-1]
	}
	curDiff := array[1]
	curCount := 0
	for i := 2; i < len(A); i++ {
		if array[i] == curDiff {
			curCount++
			count += curCount
		} else {
			curCount = 0
			curDiff = array[i]
		}
	}
	return count
}
