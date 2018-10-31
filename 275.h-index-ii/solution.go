package leetcode

func hIndex(citations []int) int {
	res := 0
	for i := 0; i < len(citations); i++ {
		curH := min(len(citations)-i, citations[i])
		res = max(res, curH)
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
