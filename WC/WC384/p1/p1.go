package p1

func modifiedMatrix(matrix [][]int) [][]int {
	m := len(matrix)
	n := len(matrix[0])
	for j := 0; j < n; j++ {
		maxJ := matrix[0][j]
		flg := false
		for i := 0; i < m; i++ {
			if matrix[i][j] == -1 && !flg {
				flg = true
			}
			if maxJ < matrix[i][j] {
				maxJ = matrix[i][j]
			}
		}
		if flg {
			for i := 0; i < m; i++ {
				if matrix[i][j] == -1 {
					matrix[i][j] = maxJ
				}
			}
		}
	}
	return matrix
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -1 * x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs64(x int64) int64 {
	if x >= 0 {
		return x
	}
	return -1 * x

}

func min64(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func max64(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}
