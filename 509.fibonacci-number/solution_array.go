package leetcode

// array
func fib_3(N int) int {
	if N < 2 {
		return N
	}
	array := make([]int, N)
	array[0] = 1
	array[1] = 1
	for i := 2; i < N; i++ {
		array[i] = array[i-1] + array[i-2]
	}
	return array[N-1]
}
