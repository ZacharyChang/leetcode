package leetcode

// Two Pointers
func maxArea(height []int) int {
	length := len(height)
	l, r := 0, length-1
	maxRes := 0
	for l < r {
		minHeight := min(height[l], height[r]) * (r - l)
		if minHeight > maxRes {
			maxRes = minHeight
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return maxRes
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
