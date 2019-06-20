package leetcode

// Hash Table
// TODO: Two Pointers
func findPairs(nums []int, k int) int {
	if k < 0 {
		return 0
	}
	count := 0
	numMap := make(map[int]int, 0)
	for _, v := range nums {
		numMap[v]++
	}
	for key, value := range numMap {
		if k == 0 {
			if value > 1 {
				// count as twice
				// the result will be same as the duplicate numbers when k equals to 0
				count += 2
			}
		} else {
			if numMap[key-k] > 0 {
				count++
			}
			if numMap[key+k] > 0 {
				count++
			}
		}
	}
	return count / 2
}
