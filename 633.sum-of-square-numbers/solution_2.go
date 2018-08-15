package leetcode

func judgeSquareSum_2(c int) bool {
	for i := 0; i < (c/2 + 1); i++ {
		if c-i*i < 0 {
			return false
		}
		if isSqure(c - i*i) {
			return true
		}
	}
	return false
}

func isSqure(num int) bool {
	i := 0
	j := num/2 + 1
	for i <= j {
		mid := (i + j) / 2
		sq := mid * mid
		if sq == num {
			return true
		} else if sq < num {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return false
}
