package leetcode

import (
	"math"
)

func arrangeCoins(n int) int {
	return int(math.Sqrt(float64(n)*2+0.25) - 0.5)
}
