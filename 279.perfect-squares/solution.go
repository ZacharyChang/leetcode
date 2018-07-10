package leetcode

func numSquares(n int) int {
	// optimize performance code
	for n%4 == 0 {
		n = n / 4
	}
	if n%8 == 7 {
		return 4
	}
	array := make([]int, n+1)
	for i := 0; i <= n; i++ {
		for j := 1; i+j*j <= n; j++ {
			if array[i+j*j] == 0 {
				array[i+j*j] = array[i] + 1
			} else if array[i]+1 < array[i+j*j] {
				array[i+j*j] = array[i] + 1
			}
		}
	}
	return array[n]
}
