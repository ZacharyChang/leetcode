package leetcode

func searchInsert_brute_force(nums []int, target int) int {
	for i, value := range nums {
		if value >= target {
			return i
		}
	}
	return len(nums)
}
