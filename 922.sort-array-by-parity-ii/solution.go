package leetcode

func sortArrayByParityII(A []int) []int {
	start, end := 0, len(A)-1
	for start < len(A) && end >= 0 {
		if !isOdd(A[start]) {
			start += 2
			continue
		}
		if isOdd(A[end]) {
			end -= 2
			continue
		}
		A[start], A[end] = A[end], A[start]
		start, end = start+2, end-2
	}
	return A
}

func isOdd(a int) bool {
	return a%2 == 1
}
