package leetcode

// plusOne returns the new array that old array plus one in the end
func plusOne(digits []int) []int {
	if digits == nil || len(digits) == 0 {
		return []int{}
	}
	digits[len(digits)-1]++
	for i := len(digits) - 1; i >= 1; i-- {
		if digits[i] > 9 {
			digits[i] = 0
			digits[i-1]++
		} else {
			return digits
		}
	}
	if digits[0] > 9 {
		return append([]int{1, 0}, digits[1:]...)
	}
	return digits
}
