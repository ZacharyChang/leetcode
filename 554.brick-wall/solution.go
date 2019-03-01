package leetcode

func leastBricks(wall [][]int) int {
	brickMap := make(map[int]int, 0)
	maxMatch := 0
	for i := 0; i < len(wall); i++ {
		curWidth := 0
		for j := 0; j < len(wall[i])-1; j++ {
			curWidth += wall[i][j]
			brickMap[curWidth]++
		}
	}
	for _, v := range brickMap {
		if v > maxMatch {
			maxMatch = v
		}
	}
	return len(wall) - maxMatch
}
