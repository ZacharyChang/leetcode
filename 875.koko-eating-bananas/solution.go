package leetcode

import (
	"sort"
)

func minEatingSpeed(piles []int, H int) int {
	sort.Ints(piles)
	start, end := 1, piles[len(piles)-1]
	for start < end {
		mid := (start + end) / 2
		if check(&piles, H, mid) {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return end
}

func check(pilesPtr *[]int, H int, speed int) bool {
	totalHour := 0
	for _, v := range *pilesPtr {
		if v%speed > 0 {
			totalHour += v/speed + 1
		} else {
			totalHour += v / speed
		}
	}
	return totalHour <= H
}
