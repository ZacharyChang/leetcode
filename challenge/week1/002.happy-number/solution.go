package challenge

import "fmt"

func isHappy(n int) bool {
	return helper(n, make(map[int]bool))
}

func helper(n int, m map[int]bool) bool {
	fmt.Printf("n: %d\n", n)
	if n == 1 {
		return true
	}
	if m[n] {
		return false
	}
	m[n] = true
	r := 0
	for ; n != 0; n /= 10 {
		r += (n % 10) * (n % 10)
	}
	return helper(r, m)
}
