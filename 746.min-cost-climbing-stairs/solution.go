package leetcode

func minCostClimbingStairs(cost []int) int {
	if cost == nil || len(cost) == 0 {
		return 0
	}
	if len(cost) == 1 {
		return cost[0]
	}
	for i := 2; i < len(cost); i++ {
		cost[i] += min(cost[i-2], cost[i-1])
	}
	return min(cost[len(cost)-1], cost[len(cost)-2])
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
