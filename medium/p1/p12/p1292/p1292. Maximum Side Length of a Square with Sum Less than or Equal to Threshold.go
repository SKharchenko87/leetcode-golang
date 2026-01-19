package p1292

func alloc(n, m int) [][]int {
	data := make([]int, n*m)
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		start := i * m
		res[i] = data[start : start+m]
	}
	return res
}

func maxSideLength(mat [][]int, threshold int) int {
	n, m := len(mat), len(mat[0])
	sums := alloc(n+1, m+1)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			sums[i+1][j+1] = sums[i+1][j] + mat[i][j]
		}
	}

	for j := 1; j <= m; j++ {
		for i := 1; i <= n; i++ {
			sums[i][j] += sums[i-1][j]
		}
	}

	maxSide := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			x := min(i, j)
			for ; x > maxSide; x-- {
				if sums[i][j]-sums[i][j-x]-(sums[i-x][j]-sums[i-x][j-x]) <= threshold && maxSide < x {
					maxSide = x
				}
			}
		}
	}

	return maxSide
}
