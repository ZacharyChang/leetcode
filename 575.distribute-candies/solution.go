package leetcode

func distributeCandies(candies []int) int {
	candyMap := make(map[int]int, len(candies))
	for _, v := range candies {
		candyMap[v]++
	}
	return min(len(candyMap), len(candies)/2)
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
