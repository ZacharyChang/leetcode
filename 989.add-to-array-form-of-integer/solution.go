package leetcode

func addToArrayForm(A []int, K int) []int {
	arr := make([]int, 0)
	for K != 0 {
		arr = append(arr, K%10)
		K = K / 10
	}
	reverse(arr)
	return merge(A, arr, 0)

}

func merge(A []int, B []int, add int) []int {
	if len(A) == 0 {
		if add == 0 {
			return B
		}
		return merge(B, []int{add}, 0)
	}
	if len(B) == 0 {
		if add == 0 {
			return A
		}
		return merge(A, []int{add}, 0)
	}
	if A[len(A)-1]+B[len(B)-1]+add <= 9 {
		return append(merge(A[:len(A)-1], B[:len(B)-1], 0), A[len(A)-1]+B[len(B)-1]+add)
	}
	return append(merge(A[:len(A)-1], B[:len(B)-1], 1), A[len(A)-1]+B[len(B)-1]+add-10)
}

func reverse(A []int) {
	for i := 0; i < len(A)-1-i; i++ {
		A[i], A[len(A)-1-i] = A[len(A)-1-i], A[i]
	}
}
