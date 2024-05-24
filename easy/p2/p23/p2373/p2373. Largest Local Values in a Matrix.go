package p2373

func largestLocal(grid [][]int) [][]int {
	n := len(grid)
	res := make([][]int, n-2)
	for i := 0; i < n-2; i++ {
		res[i] = make([]int, n-2)
		for j := 0; j < n-2; j++ {
			res[i][j] = max(
				grid[i][j], grid[i][j+1], grid[i][j+2],
				grid[i+1][j], grid[i+1][j+1], grid[i+1][j+2],
				grid[i+2][j], grid[i+2][j+1], grid[i+2][j+2],
			)
		}
	}
	return res
}
