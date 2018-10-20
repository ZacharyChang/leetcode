package leetcode

import (
	"sort"
)

func findRadius(houses []int, heaters []int) int {
	if len(houses) == 0 || len(heaters) == 0 {
		return 0
	}
	sort.Ints(heaters)
	radius := -1
	for _, v := range houses {
		radius = max(radius, minDistance(v, &heaters))
	}
	return radius
}

func minDistance(house int, heatersPointer *[]int) int {
	heaters := *heatersPointer
	if house < heaters[0] || house > heaters[len(heaters)-1] {
		return min(distance(heaters[0], house), distance(heaters[len(heaters)-1], house))
	}
	start := 0
	end := len(heaters) - 1
	for start < end-1 {
		mid := (start + end) / 2
		if house == heaters[mid] {
			return 0
		} else if house < heaters[mid] {
			end = mid
		} else {
			start = mid
		}
	}
	return min(house-heaters[start], heaters[end]-house)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func distance(a int, b int) int {
	if a-b >= 0 {
		return a - b
	}
	return b - a
}
