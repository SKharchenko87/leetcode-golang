package p0463

// ToDo benchmark
func islandPerimeter(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	d := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	res := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				res += 4
				for _, dx := range d {
					di, dj := i+dx[0], j+dx[1]
					if 0 <= di && di < rows && 0 <= dj && dj < cols && grid[di][dj] == 1 {
						res -= 1
					}
				}
			}

		}
	}
	return res
}

func islandPerimeter0(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	res := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				res += 4
				if i > 0 && grid[i-1][j] == 1 {
					res -= 1
				}
				if i < rows-1 && grid[i+1][j] == 1 {
					res -= 1
				}
				if j > 0 && grid[i][j-1] == 1 {
					res -= 1
				}
				if j < cols-1 && grid[i][j+1] == 1 {
					res -= 1
				}
			}

		}
	}
	return res
}
