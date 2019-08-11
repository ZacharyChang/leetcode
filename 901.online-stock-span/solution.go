package leetcode

// TODO: stack
type StockSpanner struct {
	prices []int
}

func Constructor() StockSpanner {
	return StockSpanner{
		prices: make([]int, 0),
	}

}

func (this *StockSpanner) Next(price int) int {
	count := 1
	for i := len(this.prices) - 1; i >= 0; i-- {
		if this.prices[i] <= price {
			count++
		} else {
			break
		}
	}
	this.prices = append(this.prices, price)
	return count
}
