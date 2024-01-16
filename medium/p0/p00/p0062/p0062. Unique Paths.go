package p0062

func uniquePaths(m int, n int) int {
	d := make([][]int, m)
	d[0] = make([]int, n)
	for j := 0; j < n; j++ {
		d[0][j] = 1
	}
	for i := 1; i < m; i++ {
		d[i] = make([]int, n)
		d[i][0] = 1
		for j := 1; j < n; j++ {
			d[i][j] = d[i-1][j] + d[i][j-1]
		}
	}
	return d[m-1][n-1]
}
