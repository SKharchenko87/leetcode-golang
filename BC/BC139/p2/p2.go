package p2

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

func findSafeWalk(grid [][]int, health int) bool {

	m, n := len(grid), len(grid[0])
	visited := make([][]int, m)
	for i := range visited {
		visited[i] = make([]int, n)
	}

	direction := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	var rec func(i, j, currentHealth int)
	rec = func(i, j, currentHealth int) {
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}
		if grid[i][j] == 1 {
			currentHealth--
			if currentHealth < 0 {
				return
			}
		}
		if currentHealth > visited[i][j] {
			visited[i][j] = currentHealth
			rec(i+direction[0][0], j+direction[0][1], currentHealth)
			rec(i+direction[1][0], j+direction[1][1], currentHealth)
			rec(i+direction[2][0], j+direction[2][1], currentHealth)
			rec(i+direction[3][0], j+direction[3][1], currentHealth)
		}
	}
	rec(0, 0, health)
	return visited[m-1][n-1] > 0
}
