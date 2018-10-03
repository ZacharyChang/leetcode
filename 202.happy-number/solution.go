package leetcode

var happyMap map[int]bool

func isHappy(n int) bool {
	happyMap = make(map[int]bool, 0)
	return helper(n)
}

func helper(n int) bool {
	if n == 1 {
		return true
	}
	if happyMap[n] == true {
		return false
	}
	happyMap[n] = true
	return helper(squareSum(n))
}

func squareSum(n int) int {
	result := 0
	for ; n > 0; n /= 10 {
		result += (n % 10) * (n % 10)
	}
	return result
}
