package leetcode

func subarraySum(nums []int, k int) int {
	sumMap := make(map[int]int, 0)
	// init with 0
	sumMap[0] = 1
	sum, count := 0, 0
	for _, num := range nums {
		sum += num
		if sumMap[sum-k] > 0 {
			count += sumMap[sum-k]
		}
		sumMap[sum]++
	}
	return count
}
