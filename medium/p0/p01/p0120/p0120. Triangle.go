package p0120

import "slices"

func minimumTotal(triangle [][]int) int {
	l := len(triangle)
	for i := 1; i < l; i++ {
		triangle[i][0] += triangle[i-1][0]
		triangle[i][i] += triangle[i-1][i-1]
		for j := 1; j < i; j++ {
			triangle[i][j] += min(triangle[i-1][j-1], triangle[i-1][j])
		}
	}
	return slices.Min(triangle[l-1])
}
