package leetcode

func countPrimes(n int) int {
	result := 0
	primeArray := make([]bool, n)
	for i := 2; i < n; i++ {
		primeArray[i] = true
	}
	for i := 2; i < n; i++ {
		if !primeArray[i] {
			continue
		}
		// number i is prime number
		result++
		for j := 2; j*i < n; j++ {
			primeArray[i*j] = false
		}
	}
	return result
}
