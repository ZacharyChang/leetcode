package leetcode

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	res := 0
	for i := 0; i < len(prices)-1; i++ {
		res += max(0, prices[i+1]-prices[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
