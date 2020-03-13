package leetcode

import (
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	res := ""
	strs := make([]string, 0)
	for _, num := range nums {
		strs = append(strs, strconv.Itoa(num))
	}
	sort.Sort(sort.Reverse(Nums(strs)))
	for _, str := range strs {
		res += str
	}
	for len(res) > 0 && res[0] == '0' {
		res = res[1:]
	}
	if len(res) == 0 {
		return "0"
	}
	return res
}

type Nums []string

func (n Nums) Len() int {
	return len(n)
}

func (n Nums) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n Nums) Less(a, b int) bool {
	return n[a]+n[b] < n[b]+n[a]
}
