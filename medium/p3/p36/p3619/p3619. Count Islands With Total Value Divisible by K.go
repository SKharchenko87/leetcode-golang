package p3619

func countIslands(grid [][]int, k int) int {
	n, m := len(grid), len(grid[0])
	var directions [][2]int = [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	var bfs func(int, int) int
	bfs = func(x, y int) int {
		if !(0 <= x && x < n && 0 <= y && y < m) {
			return 0
		}
		if grid[x][y] == 0 {
			return 0
		}
		res := grid[x][y]
		grid[x][y] = 0
		for _, d := range directions {
			res += bfs(x+d[0], y+d[1])
		}
		return res
	}
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] > 0 {
				if bfs(i, j)%k == 0 {
					res++
				}
			}
		}
	}
	return res
}
