package p2536

func rangeAddQueries(n int, queries [][]int) [][]int {
	mat := make([][]int, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]int, n)
	}
	for _, query := range queries {
		i0, j0 := query[0], query[1]
		i1, j1 := query[2], query[3]
		mat[i0][j0]++
		if j1+1 < n {
			mat[i0][j1+1]--
		}
		if i1+1 < n {
			mat[i1+1][j0]--
		}
		if i1+1 < n && j1+1 < n {
			mat[i1+1][j1+1]++
		}
	}
	var cur int
	for j := 0; j < n; j++ {
		cur += mat[0][j]
		mat[0][j] = cur
	}

	for i := 1; i < n; i++ {
		cur = 0
		for j := 0; j < n; j++ {
			cur += mat[i][j]
			mat[i][j] = cur + mat[i-1][j]
		}
	}

	return mat
}

func rangeAddQueries1(n int, queries [][]int) [][]int {
	diff := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		diff[i] = make([]int, n+1)
	}
	for _, query := range queries {
		i0, j0 := query[0], query[1]
		i1, j1 := query[2], query[3]
		diff[i0][j0]++
		diff[i0][j1+1]--
		diff[i1+1][j0]--
		diff[i1+1][j1+1]++
	}
	var cur int
	for j := 0; j < n; j++ {
		cur += diff[0][j]
		diff[0][j] = cur
	}
	diff[0] = diff[0][:n]
	for i := 1; i < n; i++ {
		cur = 0
		for j := 0; j < n; j++ {
			cur += diff[i][j]
			diff[i][j] = cur + diff[i-1][j]
		}
		diff[i] = diff[i][:n]
	}

	return diff[:n]
}

func rangeAddQueries0(n int, queries [][]int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	for _, query := range queries {
		for i := query[0]; i <= query[2]; i++ {
			for j := query[1]; j <= query[3]; j++ {
				matrix[i][j]++
			}
		}
	}
	return matrix
}
