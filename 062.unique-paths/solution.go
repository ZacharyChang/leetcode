package leetcode

func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	array := make([][]int, m)
	for i := 0; i < len(array); i++ {
		array[i] = make([]int, n)
	}
	array[len(array)-1][len(array[0])-1] = 1
	for i := len(array) - 1; i >= 0; i-- {
		for j := len(array[0]) - 1; j >= 0; j-- {
			if i+1 < len(array) {
				array[i][j] += array[i+1][j]
			}
			if j+1 < len(array[0]) {
				array[i][j] += array[i][j+1]
			}
		}
	}
	return array[0][0]
}
