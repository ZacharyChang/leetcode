package leetcode

func moveZeroes(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	//ptr1, ptr2 := 0, 1
	//for ptr1 < len(nums) && ptr2 < len(nums) {
	//	n1 := nums[ptr1]
	//	n2 := nums[ptr2]
	//
	//}
	l := len(nums)
	for i := 0; i < l; i++ {
		//for j := i + 1; j < len(nums); j++ {
		//	if nums[i] == 0 {
		//		//nums[i], nums[j] = nums[j], nums[i]
		//		copy(nums[i:len(nums)-1], nums[i+1:len(nums)])
		//		nums[len(nums)-1] = 0
		//	}
		//}
		if nums[i] == 0 {
			copy(nums[i:len(nums)-1], nums[i+1:])
			nums[len(nums)-1] = 0
			i--
			l--
		}
	}
	return nums
}

// Bubble Sort
func moveZeroes_2(nums []int) []int{

}