package leetcode

// DP
func maxProfit(prices []int) int {
	if prices == nil || len(prices) == 0 {
		return 0
	}
	buyAt := prices[0]
	result := 0
	for _, price := range prices {
		buyAt = min(buyAt, price)
		result = max(result, price-buyAt)
	}
	return result
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// Brute Force
func maxProfit_2(prices []int) int {
	if prices == nil || len(prices) == 0 {
		return 0
	}
	max := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j]-prices[i] > max {
				max = prices[j] - prices[i]
			}
		}
	}
	return max
}
