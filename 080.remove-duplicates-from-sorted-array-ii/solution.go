package leetcode

func removeDuplicates(nums []int) int {
	i := 0
	repeatCount := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			repeatCount = 0
			i++
			nums[i] = nums[j]
		} else if repeatCount == 0 {
			repeatCount++
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
