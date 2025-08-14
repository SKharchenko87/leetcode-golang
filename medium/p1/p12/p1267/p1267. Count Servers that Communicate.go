package p1267

func countServers(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	res := 0
	colCounts := make([]int, n)
	lastIndexCol := make([]int, m)
	for i := 0; i < m; i++ {
		lastIndexCol[i] = -1
	}
	for i := 0; i < m; i++ {
		tmp := 0
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				tmp++
				colCounts[j]++
				lastIndexCol[i] = j
			}
		}
		if tmp > 1 {
			res += tmp
			lastIndexCol[i] = -1
		}
	}

	for i := 0; i < m; i++ {
		if lastIndexCol[i] != -1 && colCounts[lastIndexCol[i]] > 1 {
			res++
		}
	}
	return res
}

func countServers0(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	res := 0
	rows := make([]bool, m)
	for r := 0; r < m; r++ {
		tmp := 0
		for c := 0; c < n; c++ {
			if grid[r][c] == 1 {
				tmp++
			}
		}
		rows[r] = tmp > 1
		if rows[r] {
			res += tmp
		}
	}
	for c := 0; c < n; c++ {
		tmp := 0
		use := 0
		for r := 0; r < m; r++ {
			if grid[r][c] == 1 {
				tmp++
				if rows[r] {
					use++
				}
			}
		}
		if tmp > 1 {
			res += tmp - use
		}
	}
	return res
}
