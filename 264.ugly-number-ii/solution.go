package leetcode

func nthUglyNumber(n int) int {
	if n <= 0 {
		return 0
	}
	array := make([]int, n)
	array[0] = 1
	point2 := 0
	point3 := 0
	point5 := 0
	for i := 1; i < n; i++ {
		value2 := 2 * array[point2]
		value3 := 3 * array[point3]
		value5 := 5 * array[point5]
		minValue := min(value2, value3, value5)
		// 检查重复
		if array[i-1] == minValue {
			i--
		} else {
			array[i] = minValue
		}
		if minValue == value2 {
			point2++
		} else if minValue == value3 {
			point3++
		} else if minValue == value5 {
			point5++
		}
	}
	return array[n-1]
}

func min(a int, b int, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}
