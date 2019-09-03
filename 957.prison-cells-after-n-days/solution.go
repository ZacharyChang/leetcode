package leetcode

func prisonAfterNDays(cells []int, N int) []int {
	for N = (N-1)%14 + 1; N > 0; N-- {
		newCells := make([]int, len(cells))
		for i := 1; i < len(cells)-1; i++ {
			if cells[i-1] == cells[i+1] {
				newCells[i] = 1
			} else {
				newCells[i] = 0
			}
		}
		cells = newCells
	}
	return cells
}
