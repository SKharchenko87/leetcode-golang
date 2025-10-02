package p1861

const (
	STONE    = '#'
	OBSTACLE = '*'
	EMPTY    = '.'
)

func rotateTheBox(box [][]byte) [][]byte {
	m := len(box)
	n := len(box[0])
	result := make([][]byte, n)
	for i := 0; i < m; i++ {
		index := n - 1
		for j := n - 1; j >= 0; j-- {
			if box[i][j] == OBSTACLE {
				index = j - 1
			} else if box[i][j] == STONE {
				box[i][j], box[i][index] = EMPTY, STONE
				index--
			}
		}
	}

	for j := 0; j < n; j++ {
		result[j] = make([]byte, m)
		for i := 0; i < m; i++ {
			result[j][m-i-1] = box[i][j]
		}
	}

	return result
}

func rotateTheBox2(box [][]byte) [][]byte {
	m := len(box)
	n := len(box[0])
	result := make([][]byte, n)
	for i := 0; i < m; i++ {
		index := 0
		for ; index < n && box[i][index] != STONE; index++ {
		}
		for j := index + 1; j < n; j++ {
			if box[i][j] == OBSTACLE {
				index = j + 1
			} else if box[i][j] == EMPTY {
				box[i][j], box[i][index] = STONE, EMPTY
				index++
			}
		}
	}
	for j := 0; j < n; j++ {
		result[j] = make([]byte, m)
		for i := 0; i < m; i++ {
			result[j][m-i-1] = box[i][j]
		}
	}

	return result
}

func rotateTheBox1(box [][]byte) [][]byte {
	m := len(box)
	n := len(box[0])
	result := make([][]byte, n)
	for j := 0; j < n; j++ {
		result[j] = make([]byte, m)
	}
	for i := 0; i < m; i++ {
		index := 0
		for ; index < n && box[i][index] != STONE; index++ {
			result[index][m-i-1] = box[i][index]
		}
		if index < n && box[i][index] == STONE {
			result[index][m-i-1] = box[i][index]
		}
		for j := index + 1; j < n; j++ {
			if box[i][j] == OBSTACLE {
				index = j + 1
			}
			if box[i][j] == EMPTY {
				result[index][m-i-1] = box[i][j]
				box[i][j], box[i][index] = box[i][index], box[i][j]
				index++
			}
			result[j][m-i-1] = box[i][j]
		}
	}
	return result
}

func rotateTheBox0(box [][]byte) [][]byte {
	m := len(box)
	n := len(box[0])
	result := make([][]byte, n)
	for i := 0; i < m; i++ {
		index := 0
		for ; index < n && box[i][index] != '#'; index++ {
		}
		for j := index + 1; j < n; j++ {
			if box[i][j] == '*' {
				index = j + 1
			}
			if box[i][j] == '.' {
				box[i][j], box[i][index] = box[i][index], box[i][j]
				index++
			}
		}
	}
	for j := 0; j < n; j++ {
		result[j] = make([]byte, m)
		for i := 0; i < m; i++ {
			result[j][m-i-1] = box[i][j]
		}
	}
	//for j := 0; j < n; j++ {
	//	for i := 0; i < m; i++ {
	//		fmt.Printf("%c ", result[j][i])
	//	}
	//	fmt.Println()
	//}
	return result
}
