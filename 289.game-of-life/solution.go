package leetcode

func gameOfLife(board [][]int) {
	nextBoard := make([][]int, len(board))
	for i := range nextBoard {
		nextBoard[i] = make([]int, len(board[i]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			nextBoard[i][j] = newValue(board, i, j)
		}
	}
	// copy the next board to board
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			board[i][j] = nextBoard[i][j]
		}
	}
}

func newValue(board [][]int, a int, b int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			// check the range
			if a+i >= 0 && a+i < len(board) && b+j >= 0 && b+j < len(board[a+i]) {
				count += board[a+i][b+j]
			}
		}
	}
	if board[a][b] > 0 {
		if count == 2 || count == 3 {
			return 1
		}
		return 0
	} else {
		if count == 3 {
			return 1
		}
		return 0
	}
}
