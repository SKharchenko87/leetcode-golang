package p0037

// ToDo
func solveSudoku(board [][]byte) {
	//m:=map[byte]map[byte]map[byte]struct{}{}
	flag := true
LOOP:
	for flag {
		//for _, bytes := range board {
		//	for _, b := range bytes {
		//		fmt.Printf("%c ", b)
		//	}
		//	fmt.Printf("\n")
		//}
		//fmt.Printf("\n")
		flag = false
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if board[i][j] == '.' {
					m := map[byte]struct{}{1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 6: {}, 7: {}, 8: {}, 9: {}}
					for k := 0; k < 9; k++ {
						delete(m, board[i][k]-'0')
						//if board[i][k] != '.' {
						//	delete(m, board[i][k]-'0')
						//}
					}
					for k := 0; k < 9; k++ {
						delete(m, board[k][j]-'0')
					}
					startI, startJ := i/3, j/3
					for k := 0; k < 3; k++ {
						for l := 0; l < 3; l++ {
							delete(m, board[startI*3+k][startJ*3+l]-'0')
						}
					}
					if len(m) == 1 {
						for b, _ := range m {
							board[i][j] = b + '0'
						}
						flag = true
						continue LOOP
					}
				}
			}
		}
	}

}
