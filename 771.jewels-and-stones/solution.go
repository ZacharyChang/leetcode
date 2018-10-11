package leetcode

func numJewelsInStones(J string, S string) int {
	result := 0
	stoneMap := make(map[rune]int, len(S))
	for _, v := range S {
		stoneMap[v]++
	}
	for _, v := range J {
		result += stoneMap[v]
	}
	return result
}
