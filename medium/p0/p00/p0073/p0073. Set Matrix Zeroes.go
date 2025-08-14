package p0073

func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	firstI, firstJ := getFirstZero(matrix)
	if firstI != -1 && firstJ != -1 {
		for i := firstI; i < m; i++ {
			for j := 0; j < n; j++ {
				if matrix[i][j] == 0 {
					matrix[firstI][j] = 0
					matrix[i][firstJ] = 0
				}
			}
		}
		for i := 0; i < m; i++ {
			if matrix[i][firstJ] == 0 && i != firstI {
				for j := 0; j < n; j++ {
					matrix[i][j] = 0
				}
			}
		}
		for j := 0; j < n; j++ {
			if matrix[firstI][j] == 0 && j != firstJ {
				for i := 0; i < m; i++ {
					matrix[i][j] = 0
				}
			}
		}
		for i := 0; i < m; i++ {
			matrix[i][firstJ] = 0
		}
		for j := 0; j < n; j++ {
			matrix[firstI][j] = 0
		}
	}
}

func getFirstZero(matrix [][]int) (int, int) {
	m := len(matrix)
	n := len(matrix[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

func run(matrix [][]int) [][]int {
	setZeroes(matrix)
	return matrix
}
