package leetcode

import (
	"sort"
)

func largestPerimeter(A []int) int {
	if len(A) < 3 {
		return 0
	}
	sort.Sort(sort.Reverse(sort.IntSlice(A)))
	for i := 0; i+2 < len(A); i++ {
		if A[i+1]+A[i+2] > A[i] {
			return A[i] + A[i+1] + A[i+2]
		}
	}
	return 0
}
