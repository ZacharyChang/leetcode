package leetcode

func binaryGap(N int) int {
	res := 0
	cur, last := -1, -1
	for i := 0; N != 0; N, i = N/2, i+1 {
		if N%2 == 1 {
			cur, last = i, cur
		}
		if last >= 0 {
			res = max(res, cur-last)
		}
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
