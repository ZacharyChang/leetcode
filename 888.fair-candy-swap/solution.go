package leetcode

func fairCandySwap(A []int, B []int) []int {
	sumA, sumB := 0, 0
	mapA := make(map[int]bool, 0)
	for _, v := range A {
		sumA += v
		mapA[v] = true
	}
	for _, v := range B {
		sumB += v
	}
	if sumB > sumA {
		diff := (sumB - sumA) / 2
		for _, v := range B {
			if mapA[v-diff] {
				return []int{
					v - diff, v,
				}
			}
		}
	} else {
		diff := (sumA - sumB) / 2
		for _, v := range B {
			if mapA[v+diff] {
				return []int{
					v + diff, v,
				}
			}
		}
	}
	return []int{-1, -1}
}
