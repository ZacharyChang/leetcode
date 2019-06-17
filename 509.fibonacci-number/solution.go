package leetcode

func fib(N int) int {
	if N < 2 {
		return N
	}
	lastTwo, lastOne := 1, 1
	for i := 2; i < N; i++ {
		tmp := lastOne + lastTwo
		lastTwo = lastOne
		lastOne = tmp
	}
	return lastOne
}
