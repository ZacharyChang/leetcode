package leetcode

import (
	"sort"
)

var romanMap = map[int]string{
	1:    "I",
	4:    "IV",
	5:    "V",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
}

// Hashtable
func intToRoman_2(num int) string {
	v, exist := romanMap[num]
	if exist {
		return v
	}
	max := findMax(num)
	return romanMap[max] + intToRoman(num-max)
}

func findMax(num int) int {
	nums := make([]int, 0)
	for k, _ := range romanMap {
		nums = append(nums, k)
	}
	sort.Ints(nums)
	for i, v := range nums {
		if v > num {
			return nums[i-1]
		}
	}
	return nums[len(nums)-1]
}
