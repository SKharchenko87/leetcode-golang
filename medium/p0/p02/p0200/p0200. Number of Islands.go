package p0200

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	var dbfs func(x, y int)
	dbfs = func(x, y int) {
		if x < 0 || x >= m || y < 0 || y >= n {
			return
		}
		if grid[x][y] == '1' {
			grid[x][y] = '0'
			dbfs(x-1, y)
			dbfs(x+1, y)
			dbfs(x, y-1)
			dbfs(x, y+1)
		}
	}
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				res++
				dbfs(i, j)
			}
		}
	}
	return res
}
