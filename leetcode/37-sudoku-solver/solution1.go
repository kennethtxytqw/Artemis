package main

func solve(x int, y int, board [][]byte) bool {
	if y == len(board) {
		return true
	}
	if x == len(board) {
		return solve(0, y+1, board)
	}
	if board[x][y] != '.' {
		return solve(x+1, y, board)
	}
	for v := '1'; v <= '9'; v++ {
		if isLegal(byte(v), x, y, board) {
			board[x][y] = byte(v)
			if solve(x+1, y, board) {
				return true
			} else {
				board[x][y] = '.'
			}
		}
	}
	return false
}

func isLegal(v byte, x int, y int, board [][]byte) bool {
	subbox_x := x / 3 * 3
	subbox_y := y / 3 * 3
	for i := 0; i < len(board); i++ {
		if board[x][i] == v ||
			board[i][y] == v ||
			board[subbox_x+i%3][subbox_y+i/3] == v {
			return false
		}
	}
	return true
}

func solveSudoku(board [][]byte) {
	solve(0, 0, board)
}
