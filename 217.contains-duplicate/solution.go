package leetcode

func containsDuplicate(nums []int) bool {
	numsMap := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		if numsMap[nums[i]] >= 1 {
			return true
		}
		numsMap[nums[i]] = 1

	}
	return false
}
