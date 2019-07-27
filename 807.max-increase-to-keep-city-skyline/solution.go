package leetcode

func maxIncreaseKeepingSkyline(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	res := 0
	maxHeightInRow := make([]int, len(grid))
	maxHeightInColumn := make([]int, len(grid[0]))
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			maxHeightInRow[i] = max(maxHeightInRow[i], grid[i][j])
			maxHeightInColumn[j] = max(maxHeightInColumn[j], grid[i][j])
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			res += min(maxHeightInRow[i], maxHeightInColumn[j]) - grid[i][j]
		}
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
