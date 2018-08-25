package leetcode

func mySqrt(x int) int {
	// x/2+2为保证x=1时正确
	// 因为 (x/2+2)^2 > (x/2+1)^2 = x^2/4+x+1 > x
	// 故sqrt(x) < x/2+1
	start, end := 0, x/2+2
	for start < end-1 {
		mid := (start + end) / 2
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			start = mid
		} else {
			end = mid
		}
	}
	return start
}
