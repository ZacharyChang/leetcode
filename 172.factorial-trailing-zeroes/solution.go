package leetcode

func trailingZeroes(n int) int {
	result := 0
	for i := n; i > 0; i /= 5 {
		result += i / 5
	}
	return result
}
