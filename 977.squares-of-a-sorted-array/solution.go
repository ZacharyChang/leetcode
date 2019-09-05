package leetcode

import "sort"

// TODO: two pointers
func sortedSquares(A []int) []int {
	res := make([]int, 0)
	for i, v := range A {
		if v < 0 {
			A[i] = -A[i]
		} else {
			break
		}
	}
	sort.Ints(A)
	for _, v := range A {
		res = append(res, v*v)
	}
	return res
}
