package p3459

type diagonal struct {
	upLeft    int
	upRight   int
	downRight int
	downLeft  int
}

func lenOfVDiagonal(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	diagonals := make([][]diagonal, m)
	for i := 0; i < m; i++ {
		diagonals[i] = make([]diagonal, n)
	}
	// up
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ni := i - 1
			nj := j - 1
			if 0 <= ni && ni < m && 0 <= nj && nj < n && abs(grid[i][j]-grid[ni][nj]) == 2 {
				diagonals[i][j].upLeft = diagonals[ni][nj].upLeft + 1
			} else {
				diagonals[i][j].upLeft = 1
			}
			ni = i - 1
			nj = j + 1
			if 0 <= ni && ni < m && 0 <= nj && nj < n && abs(grid[i][j]-grid[ni][nj]) == 2 {
				diagonals[i][j].upRight = diagonals[ni][nj].upRight + 1
			} else {
				diagonals[i][j].upRight = 1
			}

		}
	}

	// down
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			ni := i + 1
			nj := j - 1
			if 0 <= ni && ni < m && 0 <= nj && nj < n && abs(grid[i][j]-grid[ni][nj]) == 2 {
				diagonals[i][j].downLeft = diagonals[ni][nj].downLeft + 1
			} else {
				diagonals[i][j].downLeft = 1
			}
			ni = i + 1
			nj = j + 1
			if 0 <= ni && ni < m && 0 <= nj && nj < n && abs(grid[i][j]-grid[ni][nj]) == 2 {
				diagonals[i][j].downRight = diagonals[ni][nj].downRight + 1
			} else {
				diagonals[i][j].downRight = 1
			}

		}
	}

	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				res = max(res, 1)
				ni, nj := i-1, j-1
				if 0 <= ni && ni < m && 0 <= nj && nj < n && grid[ni][nj] == 2 {
					for d := 0; d < diagonals[ni][nj].upLeft; d++ {
						res = max(res, d+diagonals[ni-d][nj-d].upRight+1)
					}
				}

				ni, nj = i-1, j+1
				if 0 <= ni && ni < m && 0 <= nj && nj < n && grid[ni][nj] == 2 {
					for d := 0; d < diagonals[ni][nj].upRight; d++ {
						res = max(res, d+diagonals[ni-d][nj+d].downRight+1)
					}
				}

				ni, nj = i+1, j-1
				if 0 <= ni && ni < m && 0 <= nj && nj < n && grid[ni][nj] == 2 {
					for d := 0; d < diagonals[ni][nj].downLeft; d++ {
						res = max(res, d+diagonals[ni+d][nj-d].upLeft+1)
					}
				}

				ni, nj = i+1, j+1
				if 0 <= ni && ni < m && 0 <= nj && nj < n && grid[ni][nj] == 2 {
					for d := 0; d < diagonals[ni][nj].downRight; d++ {
						res = max(res, d+diagonals[ni+d][nj+d].downLeft+1)
					}
				}

			}
		}
	}
	return res

}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -1 * x
}
