package leetcode

func findLUSlength(a string, b string) int {
	if len(a) != len(b) {
		return max(len(a), len(b))
	}
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] != b[i] {
			return i + 1
		}
	}
	return -1
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
