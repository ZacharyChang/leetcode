package leetcode

func validMountainArray(A []int) bool {
	if len(A) < 3 {
		return false
	}
	start, end := 0, len(A)-1
	for start <= end && start+1 < len(A) && end-1 >= 0 {
		if A[start] < A[start+1] {
			start++
			continue
		}
		if A[end-1] > A[end] {
			end--
			continue
		}
		return start == end
	}
	return false
}
