package leetcode

// Brute Force
// TODO: Two Pointers
func maxArea(height []int) int {
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

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
