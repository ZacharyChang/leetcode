package leetcode

func sortedSquares(A []int) []int {
	res := make([]int, 0)
	idx := binarySearch(A, 0)
	ptr0, ptr1 := idx-1, idx
	for ptr0 >= 0 && ptr1 < len(A) {
		if -A[ptr0] < A[ptr1] {
			res = append(res, A[ptr0]*A[ptr0])
			ptr0--
		} else {
			res = append(res, A[ptr1]*A[ptr1])
			ptr1++
		}
	}
	if ptr0 <= 0 {
		for ; ptr1 < len(A); ptr1++ {
			res = append(res, A[ptr1]*A[ptr1])
		}
	}
	if ptr1 >= len(A) {
		for ; ptr0 >= 0; ptr0-- {
			res = append(res, A[ptr0]*A[ptr0])
		}
	}
	return res
}

func binarySearch(A []int, target int) int {
	start, end := 0, len(A)-1
	if A[start] == target {
		return start
	}
	if A[end] == target {
		return end
	}
	for start < end-1 {
		mid := (start + end) / 2
		if A[mid] == target {
			return mid
		} else if A[mid] < target {
			start = mid
		} else {
			end = mid
		}
	}
	return end
}
