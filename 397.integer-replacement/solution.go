package leetcode

func integerReplacement(n int) int {
	if n == 1 {
		return 0
	}
	if n%2 == 0 {
		return integerReplacement(n/2) + 1
	}
	return min(integerReplacement(n+1), integerReplacement(n-1)) + 1
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
