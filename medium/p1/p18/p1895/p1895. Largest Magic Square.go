package p1895

func largestMagicSquare(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	rows := make([][]int, n+1)
	rows[0] = make([]int, m+1)
	cols := make([][]int, n+1)
	cols[0] = make([]int, m+1)
	for i := 1; i <= n; i++ {
		rows[i] = make([]int, m+1)
		cols[i] = make([]int, m+1)
		for j := 1; j <= m; j++ {
			rows[i][j] = rows[i][j-1] + grid[i-1][j-1]
			cols[i][j] = cols[i-1][j] + grid[i-1][j-1]
		}
	}

	diagMain := make([][]int, n+1)
	diagMain[0] = make([]int, m+1)
	for i := 1; i <= n; i++ {
		diagMain[i] = make([]int, m+1)
		for j := 1; j <= m; j++ {
			diagMain[i][j] = diagMain[i-1][j-1] + grid[i-1][j-1]
		}
	}
	diagSub := make([][]int, n+1)
	diagSub[0] = make([]int, m+2)
	for i := 1; i <= n; i++ {
		diagSub[i] = make([]int, m+2)
		for j := 1; j <= m; j++ {
			diagSub[i][j] = diagSub[i-1][j+1] + grid[i-1][j-1]
		}
	}

	var check = func(i, j, k int) bool {
		diagonalMain := diagMain[i+k][j+k] - diagMain[i][j]
		diagonalSub := diagSub[i+k][j+1] - diagSub[i][j+k+1]
		if diagonalSub != diagonalMain {
			return false
		}
		for index := 0; index < k; index++ {
			row := rows[i+1+index][j+k] - rows[i+1+index][j]
			if row != diagonalMain {
				return false
			}
			col := cols[i+k][j+1+index] - cols[i][j+1+index]
			if col != diagonalMain {
				return false
			}
		}
		return true
	}
	res := 1
	for i := 0; i < n-res; i++ {
		for j := 0; j < m-res; j++ {
			for y := min(n-i, m-j, n, m); y > res; y-- {
				if check(i, j, y) {
					res = y
				}
			}
		}
	}
	return res
}
