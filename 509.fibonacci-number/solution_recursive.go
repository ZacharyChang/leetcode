package leetcode

// recursive
func fib_2(N int) int {
	if N < 2 {
		return N
	}
	return fib_2(N-1) + fib_2(N-2)
}
