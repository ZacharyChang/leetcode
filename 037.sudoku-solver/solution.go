package leetcode

func solveSudoku(board [][]byte) {
	board = findSudokuSolution(board)
}

func findSudokuSolution(board [][]byte) [][]byte {
	if isSolved(board) {
		return board
	}

	x, y := findFirstUnknown(board)
	poss := getPossible(board, x, y)

	for _, v := range poss {
		board[x][y] = v
		if res := findSudokuSolution(board); res != nil {
			return res
		}
		// important! backtrack
		board[x][y] = '.'
	}
	// no solution
	return nil
}

func findFirstUnknown(board [][]byte) (int, int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				return i, j
			}
		}
	}
	return -1, -1
}

func isSolved(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '.' {
				return false
			}
		}
	}
	return true
}

func getPossible(board [][]byte, row int, column int) []byte {
	res := make([]byte, 0)
	existMap := make(map[byte]bool, 0)
	for i := 0; i < 9; i++ {
		if v := board[row][i]; v != '.' {
			existMap[v] = true
		}
		if v := board[i][column]; v != '.' {
			existMap[v] = true
		}
	}
	for i := row / 3 * 3; i < row/3*3+3; i++ {
		for j := column / 3 * 3; j < column/3*3+3; j++ {
			if v := board[i][j]; v != '.' {
				existMap[v] = true
			}
		}
	}
	for i := '1'; i <= '9'; i++ {
		if !existMap[byte(i)] {
			res = append(res, byte(i))
		}
	}
	return res
}
