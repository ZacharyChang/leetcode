package leetcode

// hashtable
// TODO: bit manipulation
func singleNumber(nums []int) int {
	numMap := make(map[int]int, 0)
	for _, v := range nums {
		numMap[v]++
	}
	for k, v := range numMap {
		if v == 1 {
			return k
		}
	}
	return 0
}
