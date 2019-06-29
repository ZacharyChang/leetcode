package leetcode

func checkPossibility(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			tmp := nums[i]
			nums[i] = nums[i-1]
			if checkNonDecreasing(&nums) {
				return true
			}
			nums[i-1] = tmp
			nums[i] = tmp
			return checkNonDecreasing(&nums)
		}
	}
	return true
}

func checkNonDecreasing(numsPtr *[]int) bool {
	nums := *numsPtr
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			return false
		}
	}
	return true
}
