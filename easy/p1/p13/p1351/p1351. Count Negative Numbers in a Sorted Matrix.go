package p1351

func countNegatives(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	res := 0
	j := n - 1
	for i := 0; i < m; i++ {
		for ; j >= 0 && grid[i][j] < 0; j-- {
		}
		res += n - j - 1
	}
	return res
}
