package p0931

import "math"

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	if n == 1 {
		return matrix[0][0]
	}
	var getPrev = func(i, j int) int {
		if i < 0 || j < 0 || i >= n || j >= n {
			return math.MaxInt
		}
		return matrix[i][j]
	}
	i := 1
	for ; i < n-1; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j] = matrix[i][j] + min(getPrev(i-1, j-1), getPrev(i-1, j), getPrev(i-1, j+1))
		}
	}
	res := math.MaxInt
	for j := 0; j < n; j++ {
		matrix[i][j] = matrix[i][j] + min(getPrev(i-1, j-1), getPrev(i-1, j), getPrev(i-1, j+1))
		if res > matrix[i][j] {
			res = matrix[i][j]
		}
	}
	return res
}
