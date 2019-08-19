package leetcode

func hasGroupsSizeX(deck []int) bool {
	deckMap := make(map[int]int, 0)
	for _, v := range deck {
		deckMap[v]++
	}
	x := 0
	for _, v := range deckMap {
		if x == 0 {
			x = v
		}
		if x = maxCommonDivisor(v, x); x == -1 {
			return false
		}
	}
	return x >= 2
}

func maxCommonDivisor(a int, b int) int {
	for i := min(a, b); i >= 2; i-- {
		if a%i == 0 && b%i == 0 {
			return i
		}
	}
	return -1
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
