package leetcode

func isValidSudoku(board [][]byte) bool {
	if len(board) != 9 || len(board[0]) != 9 {
		return false
	}
	// check every row
	for _, v := range board {
		if isDuplicate(v) {
			return false
		}
	}
	// check every column
	columns := make([][]byte, 9)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			columns[j] = append(columns[j], board[i][j])
		}
	}
	for _, v := range columns {
		if isDuplicate(v) {
			return false
		}
	}
	// check every part
	parts := make([][]byte, 9)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			pos := i/3*3 + j/3
			parts[pos] = append(parts[pos], board[i][j])
		}
	}
	for _, v := range parts {
		if isDuplicate(v) {
			return false
		}
	}
	return true
}

func isDuplicate(array []byte) bool {
	countMap := make(map[byte]bool, 0)
	for _, v := range array {
		if v == '.' {
			continue
		}
		if countMap[v] {
			return true
		}
		countMap[v] = true
	}
	return false
}
