package leetcode

// Brute Force
func maxArea_2(height []int) int {
	length := len(height)
	max := 0
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			num := min(height[i], height[j]) * (j - i)
			if num > max {
				max = num
			}
		}
	}
	return max
}
