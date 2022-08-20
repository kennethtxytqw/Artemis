package main

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if !isValid(i, j, board) {
				return false
			}
		}
	}
	return true
}

func isValid(i int, j int, board [][]byte) bool {
	cellValue := board[i][j]
	if cellValue == '.' {
		return true
	}
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if board[x+i/3*3][y+j/3*3] == cellValue && !(y+j/3*3 == j && x+i/3*3 == i) {
				return false
			}
		}
	}

	for x := 0; x < 9; x++ {
		if (i != x && board[x][j] == cellValue) || (j != x && board[i][x] == cellValue) {
			return false
		}
	}
	return true
}
