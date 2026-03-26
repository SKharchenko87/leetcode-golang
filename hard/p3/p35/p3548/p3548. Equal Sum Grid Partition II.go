package p3548

// ToDo
func canPartitionGrid(grid [][]int) bool {
	var total int64 = 0
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			total += int64(grid[i][j])
		}
	}
	for k := 0; k < 4; k++ {
		exist := make(map[int64]bool)
		exist[0] = true
		var sum int64 = 0
		m, n = len(grid), len(grid[0])
		if m < 2 {
			grid = rotation(grid)
			continue
		}
		if n == 1 {
			for i := 0; i < m-1; i++ {
				sum += int64(grid[i][0])
				tag := sum*2 - total
				if tag == 0 || tag == int64(grid[0][0]) || tag == int64(grid[i][0]) {
					return true
				}
			}
			grid = rotation(grid)
			continue
		}
		for i := 0; i < m-1; i++ {
			for j := 0; j < n; j++ {
				exist[int64(grid[i][j])] = true
				sum += int64(grid[i][j])
			}
			tag := sum*2 - total
			if i == 0 {
				if tag == 0 || tag == int64(grid[0][0]) || tag == int64(grid[0][n-1]) {
					return true
				}
				continue
			}
			if exist[tag] {
				return true
			}
		}
		grid = rotation(grid)
	}
	return false
}

func rotation(grid [][]int) [][]int {
	m, n := len(grid), len(grid[0])
	tmp := make([][]int, n)
	for i := range tmp {
		tmp[i] = make([]int, m)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			tmp[j][m-1-i] = grid[i][j]
		}
	}
	return tmp
}
