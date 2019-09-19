package leetcode

import (
	"sort"
)

func lastStoneWeight(stones []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(stones)))
	if len(stones) == 0 {
		return 0
	}
	if len(stones) == 1 {
		return stones[0]
	}
	if stones[0] == stones[1] {
		return lastStoneWeight(stones[2:])
	}
	stones[1] = stones[0] - stones[1]
	return lastStoneWeight(stones[1:])
}
