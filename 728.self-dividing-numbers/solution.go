package leetcode

func selfDividingNumbers(left int, right int) []int {
	result := make([]int, 0)
	for i := left; i <= right; i++ {
		if isSelfDividingNumber(i) {
			result = append(result, i)
		}
	}
	return result
}

func isSelfDividingNumber(num int) bool {
	for i := 1; i <= num; i *= 10 {
		cur := num / i % 10
		if cur == 0 {
			return false
		}
		if num%cur != 0 {
			return false
		}
	}
	return true
}
