package leetcode

func countElements(arr []int) int {
	m := make(map[int]int)
	for _, v := range arr {
		m[v]++
	}
	res := 0
	for k := range m {
		if m[k+1] > 0 {
			res += m[k]
		}
	}
	return res
}
