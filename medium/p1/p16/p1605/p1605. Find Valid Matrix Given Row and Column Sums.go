package p1605

func restoreMatrix(rowSum []int, colSum []int) [][]int {
	j := 0
	n, m := len(rowSum), len(colSum)
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		if res[i] == nil {
			res[i] = make([]int, m)
		}
		minSum := min(rowSum[i], colSum[j])
		res[i][j] = minSum
		colSum[j] -= minSum
		rowSum[i] -= minSum
		if rowSum[i] != 0 {
			j++
			i--
		}
	}
	return res
}

func restoreMatrix0(rowSum []int, colSum []int) [][]int {
	i, j := 0, 0
	n, m := len(rowSum), len(colSum)
	res := make([][]int, n)
	for i < n && j < m {
		if res[i] == nil {
			res[i] = make([]int, m)
		}
		if rowSum[i] < colSum[j] {
			res[i][j] = rowSum[i]
			colSum[j] -= rowSum[i]
			rowSum[i] = 0
			i++
		} else {
			res[i][j] = colSum[j]
			rowSum[i] -= colSum[j]
			colSum[j] = 0
			j++
		}
	}
	for i < n {
		if res[i] == nil {
			res[i] = make([]int, m)
		}
		i++
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
