package p0036

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		if !isValidRow(board[i]) {
			return false
		}
	}
	for i := 0; i < 9; i++ {
		if !isValidColumn(board, i) {
			return false
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if !isValidSubBoxes(board, i*3, j*3) {
				return false
			}
		}
	}
	return true
}

func isValidRow(board []byte) bool {
	tmp := make([]bool, 9)
	for _, b := range board {
		if b != '.' {
			if !tmp[b-'1'] {
				tmp[b-'1'] = true
			} else {
				return false
			}
		}
	}
	return true
}

func isValidColumn(board [][]byte, startJ int) bool {
	tmp := make([]bool, 9)
	for i := 0; i < 9; i++ {
		if board[i][startJ] != '.' {
			if !tmp[board[i][startJ]-'1'] {
				tmp[board[i][startJ]-'1'] = true
			} else {
				return false
			}
		}
	}
	return true
}
func isValidSubBoxes(board [][]byte, startI, startJ int) bool {
	tmp := make([]bool, 9)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i+startI][j+startJ] != '.' {
				if !tmp[board[i+startI][j+startJ]-'1'] {
					tmp[board[i+startI][j+startJ]-'1'] = true
				} else {
					return false
				}
			}
		}
	}
	return true
}
