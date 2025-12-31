package p0840

func numMagicSquaresInside(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	cnt := 0
	for i := 1; i < n-1; i++ {
		for j := 1; j < m-1; j++ {
			f := 0
			f |= 1 << grid[i-1][j-1]
			f |= 1 << grid[i-1][j]
			f |= 1 << grid[i-1][j+1]

			f |= 1 << grid[i][j-1]
			f |= 1 << grid[i][j]
			f |= 1 << grid[i][j+1]

			f |= 1 << grid[i+1][j-1]
			f |= 1 << grid[i+1][j]
			f |= 1 << grid[i+1][j+1]
			if f != 0b1111111110 {
				continue
			}
			if grid[i][j] != 5 {
				continue
			}
			row0 := grid[i-1][j-1] + grid[i-1][j] + grid[i-1][j+1]
			if row0 != 15 {
				continue
			}
			row1 := grid[i][j-1] + grid[i][j+1]
			if row1 != 10 {
				continue
			}
			row2 := grid[i+1][j-1] + grid[i+1][j] + grid[i+1][j+1]
			if row2 != 15 {
				continue
			}

			col0 := grid[i-1][j-1] + grid[i][j-1] + grid[i+1][j-1]
			if col0 != 15 {
				continue
			}
			col1 := grid[i-1][j] + grid[i+1][j]
			if col1 != 10 {
				continue
			}
			col2 := grid[i-1][j+1] + grid[i][j+1] + grid[i+1][j+1]
			if col2 != 15 {
				continue
			}

			dig0 := grid[i-1][j-1] + grid[i+1][j+1]
			if dig0 != 10 {
				continue
			}
			dig1 := grid[i-1][j+1] + grid[i+1][j-1]
			if dig1 != 10 {
				continue
			}

			cnt++
		}
	}
	return cnt
}
