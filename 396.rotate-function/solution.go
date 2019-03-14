package leetcode

// F(0) = 0*A + 1*B + 2*C + 3*D
// F(1) = 0*D + 1*A + 2*B + 3*C
// F(2) = 0*C + 1*D + 2*A + 3*B
// F(3) = 0*B + 1*C + 2*D + 3*A
// ...
// F(i) = F(i-1) + sum - n*A[n-i]
func maxRotateFunction(A []int) int {
	n := len(A)
	if n <= 1 {
		return 0
	}
	sum := A[0]
	temp := 0
	for i := 1; i < n; i++ {
		sum += A[i]
		temp += i * A[i]
	}
	max := temp
	for i := 1; i < n; i++ {
		temp = temp + sum - n*A[n-i]
		if temp > max {
			max = temp
		}
	}
	return max
}
