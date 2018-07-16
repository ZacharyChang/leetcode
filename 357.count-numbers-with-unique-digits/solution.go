package leetcode

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 10
	}
	result := make([]int, n+1)
	result[0] = 1
	result[1] = result[0] + 9
	dpTmp := 9
	for i := 2; i <= n; i++ {
		dpTmp = dpTmp * (10 - (i - 1))
		result[i] = result[i-1] + dpTmp
	}
	return result[n]
}
