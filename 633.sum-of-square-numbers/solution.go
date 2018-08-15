package leetcode

func judgeSquareSum(c int) bool {
	i := 0
	j := square(c)
	for i <= j {
		sum := i*i + j*j
		if sum == c {
			return true
		} else if sum > c {
			j--
		} else {
			i++
		}
	}
	return false
}

func square(num int) int {
	i := 0
	j := num/2 + 1
	for i <= j {
		mid := (i + j) / 2
		sq := mid * mid
		if sq == num {
			return sq
		} else if sq < num {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return j
}
