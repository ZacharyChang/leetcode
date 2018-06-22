package leetcode

// count sort and overwrite
// time complexity: O(n)
// space complexity: O(1)
func sortColors(nums []int) {
	countZero, countOne, countTwo := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		switch nums[i] {
		case 0:
			countZero++
		case 1:
			countOne++
		case 2:
			countTwo++
		}
	}
	for i := 0; i < len(nums); i++ {
		if countZero > 0 {
			nums[i] = 0
			countZero--
		} else if countOne > 0 {
			nums[i] = 1
			countOne--
		} else {
			nums[i] = 2
		}
	}
}
