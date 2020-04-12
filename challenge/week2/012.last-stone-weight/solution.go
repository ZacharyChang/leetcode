package leetcode

import "sort"

func lastStoneWeight(stones []int) int {
	if len(stones) == 0 {
		return 0
	}
	if len(stones) == 1 {
		return stones[0]
	}
	sort.Ints(stones)
	x, y := stones[len(stones)-2], stones[len(stones)-1]
	if x == y {
		stones = stones[:len(stones)-2]
	} else {
		stones = append(stones[:len(stones)-2], y-x)
	}
	return lastStoneWeight(stones)
}
