package leetcode

func isPerfectSquare(num int) bool {
	if num == 0 || num == 1 {
		return true
	}
	start, end := 0, num/2+1
	for start < end-1 {
		mid := (start + end) / 2
		if mid*mid == num {
			return true
		} else if mid*mid < num {
			start = mid
		} else {
			end = mid
		}
	}
	return false
}
