package p2684

func maxMoves(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		grid[i][0] *= -1
	}
	var cnt bool
	for j := 1; j < n; j++ {
		cnt = true
		for i := 0; i < m; i++ {
			if i-1 >= 0 && grid[i-1][j-1] < 0 && grid[i-1][j-1] > -grid[i][j] ||
				grid[i][j-1] < 0 && grid[i][j-1] > -grid[i][j] ||
				i+1 < m && grid[i+1][j-1] < 0 && grid[i+1][j-1] > -grid[i][j] {
				grid[i][j] = -grid[i][j]
				cnt = false
			}
		}
		if cnt {
			return j - 1
		}
	}
	return n - 1
}

func maxMoves0(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		grid[i][0] *= -1
	}
	var cnt int
	for j := 1; j < n; j++ {
		cnt = 0
		for i := 0; i < m; i++ {
			c := -grid[i][j]
			if i-1 >= 0 && grid[i-1][j-1] < 0 && grid[i-1][j-1] > c ||
				grid[i][j-1] < 0 && grid[i][j-1] > c ||
				i+1 < m && grid[i+1][j-1] < 0 && grid[i+1][j-1] > c {
				grid[i][j] = c
				cnt++
			}
		}
		if cnt == 0 {
			return j - 1
		}
	}
	return n - 1

}
