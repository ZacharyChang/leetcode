package leetcode

func findLUSlength(a string, b string) int {
	if len(a) != len(b) {
		return max(len(a), len(b))
	}
	if a == b {
		return -1
	}
	return len(a)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
