package leetcode

func containsNearbyDuplicate(nums []int, k int) bool {
	numsMap := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		val, exists := numsMap[nums[i]]
		if exists && i-val <= k {
			return true
		}
		numsMap[nums[i]] = i
	}
	return false
}
