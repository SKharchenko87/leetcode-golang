package p3564

func canPartitionGrid(grid [][]int) bool {
	m, n := len(grid), len(grid[0])
	rows := make([]int, m)
	cols := make([]int, n)
	sum := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sum += grid[i][j]
			rows[i] += grid[i][j]
			cols[j] += grid[i][j]
		}
	}
	if sum%2 == 1 {
		return false
	}
	sum /= 2
	sumRow := 0
	for i := 0; i < m && sumRow < sum; i++ {
		sumRow += rows[i]
		if sumRow == sum {
			return true
		}
	}
	sumCol := 0
	for j := 0; j < n; j++ {
		sumCol += cols[j]
		if sumCol == sum {
			return true
		}
	}

	return false

}
