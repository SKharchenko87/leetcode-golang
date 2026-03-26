package p3567

import (
	"math"
	"sort"
)

func minAbsDiff(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	res := make([][]int, m-k+1)
	for i := range res {
		res[i] = make([]int, n-k+1)
	}
	arr := make([]int, k*k)
	for i := 0; i < m-k+1; i++ {
		for j := 0; j < n-k+1; j++ {
			index := 0
			for i0 := 0; i0 < k; i0++ {
				for j0 := 0; j0 < k; j0++ {
					arr[index] = grid[i+i0][j+j0]
					index++
				}
			}
			sort.Ints(arr)
			x := math.MaxInt
			for a := 1; a < k*k; a++ {
				y := min(abs(arr[a-1]-arr[a]), abs(arr[a]-arr[a-1]))
				if y > 0 && y < x {
					x = y
				}
			}
			if x == math.MaxInt {
				res[i][j] = 0
			} else {
				res[i][j] = x
			}
		}
	}

	return res
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -1 * x
}
