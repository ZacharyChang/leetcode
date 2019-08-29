package leetcode

func powerfulIntegers(x int, y int, bound int) []int {
	return helper(powerArray(x, bound), powerArray(y, bound), bound)
}

func helper(arr1 []int, arr2 []int, bound int) []int {
	res := make([]int, 0)
	resMap := make(map[int]int, 0)
	for _, v1 := range arr1 {
		for _, v2 := range arr2 {
			if v1+v2 <= bound {
				resMap[v1+v2]++
			}
		}
	}
	for k := range resMap {
		res = append(res, k)
	}
	return res
}

func powerArray(x int, bound int) []int {
	res := make([]int, 0)
	cur := 1
	res = append(res, cur)
	if x == 1 {
		return res
	}
	for ; cur <= bound; cur = cur * x {
		res = append(res, cur)
	}
	return res
}
