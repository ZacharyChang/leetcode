package leetcode

func bitwiseComplement(N int) int {
	if N == 0 {
		return 1
	}
	fact := 1
	for fact-1 < N {
		fact *= 2
	}
	return fact - 1 - N
}
