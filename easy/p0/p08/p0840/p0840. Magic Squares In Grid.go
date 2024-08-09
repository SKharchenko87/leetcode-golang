package p0840

import (
	"slices"
	"sort"
)

func numMagicSquaresInside(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])
	res := 0
	for i := 1; i < row-1; i++ {
		for j := 1; j < col-1; j++ {
			if
			// row
			grid[i-1][j-1]+grid[i-1][j]+grid[i-1][j+1] == 15 &&
				grid[i][j-1]+grid[i][j]+grid[i][j+1] == 15 &&
				grid[i+1][j-1]+grid[i+1][j]+grid[i+1][j+1] == 15 &&
				// col
				grid[i-1][j-1]+grid[i][j-1]+grid[i+1][j-1] == 15 &&
				grid[i-1][j]+grid[i][j]+grid[i+1][j] == 15 &&
				grid[i-1][j+1]+grid[i][j+1]+grid[i+1][j+1] == 15 &&
				// diagonal
				grid[i-1][j-1]+grid[i][j]+grid[i+1][j+1] == 15 &&
				grid[i-1][j+1]+grid[i][j]+grid[i+1][j-1] == 15 {
				tmp := []int{grid[i-1][j-1], grid[i-1][j], grid[i-1][j+1],
					grid[i][j-1], grid[i][j], grid[i][j+1],
					grid[i+1][j-1], grid[i+1][j], grid[i+1][j+1],
				}
				sort.Ints(tmp)
				if slices.Compare(tmp, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}) == 0 {
					res++
				}
			}
		}
	}
	return res
}
