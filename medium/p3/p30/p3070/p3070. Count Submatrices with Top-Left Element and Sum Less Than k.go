package p3070

func countSubmatrices(grid [][]int, k int) int {
	res := 0
	m, n := len(grid), len(grid[0])

	sum := 0
	for j := 0; j < n; j++ {
		sum += grid[0][j]
		grid[0][j] = sum
		if grid[0][j] > k {
			n = j
			break
		} else {
			res++
		}
	}

	for i := 1; i < m; i++ {
		sum := 0
		for j := 0; j < n; j++ {
			sum += grid[i][j]
			grid[i][j] = sum + grid[i-1][j]
			if grid[i][j] > k {
				n = j
				break
			} else {
				res++
			}
		}
	}

	return res
}
