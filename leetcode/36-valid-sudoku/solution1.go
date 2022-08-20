package main

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if !isLegal(board[i][j], i, j, board) {
				return false
			}
		}
	}
	return true
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
