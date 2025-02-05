package p2658

func findMaxFish(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var rec func(i, j int) int
	rec = func(i, j int) int {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
			return 0
		}
		tmp := grid[i][j]
		grid[i][j] = 0
		return tmp + rec(i-1, j) + rec(i+1, j) + rec(i, j-1) + rec(i, j+1)
	}
	result := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 0 {
				if x := rec(i, j); result < x {
					result = x
				}
			}
		}
	}
	return result
}
