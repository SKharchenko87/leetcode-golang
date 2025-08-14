package p1568

import "fmt"

func minDays(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var recursive func(i, j, color int) int
	recursive = func(i, j, color int) int {
		if 0 <= i && i < m && 0 <= j && j < n && grid[i][j] == 1 {
			grid[i][j] = color
			res := 1
			res += recursive(i, j-1, color)
			res += recursive(i+1, j, color)
			res += recursive(i, j+1, color)
			res += recursive(i-1, j, color)
			return res
		}
		return 0
	}

	color := 2
	minCnt := 900
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if grid[r][c] == 1 {
				tmp := recursive(r, c, color)
				if tmp < minCnt {
					minCnt = tmp
				}
				color++
			}
		}
	}
	if color > 3 {
		return 0
	}
	if minCnt == 2 {
		return 2
	}
	fmt.Println(grid)
	fmt.Println(1 << 1)
	return 0
}

func getBin(i, j, m, n int, grid [][]int) int {
	res := 0
	y, x := i-1, j-1
	if 0 <= y && y <= m && 0 <= x && x < n && grid[y][x] == 1 {
		res |= 1 << 8
	}
	y, x = i-1, j
	if 0 <= y && y <= m && 0 <= x && x < n && grid[y][x] == 1 {
		res |= 1 << 7
	}
	y, x = i-1, j+1
	if 0 <= y && y <= m && 0 <= x && x < n && grid[y][x] == 1 {
		res |= 1 << 6
	}

	y, x = i, j-1
	if 0 <= y && y <= m && 0 <= x && x < n && grid[y][x] == 1 {
		res |= 1 << 5
	}
	y, x = i, j
	if 0 <= y && y <= m && 0 <= x && x < n && grid[y][x] == 1 {
		res |= 1 << 4
	}
	y, x = i, j+1
	if 0 <= y && y <= m && 0 <= x && x < n && grid[y][x] == 1 {
		res |= 1 << 3
	}

	y, x = i+1, j-1
	if 0 <= y && y <= m && 0 <= x && x < n && grid[y][x] == 1 {
		res |= 1 << 2
	}
	y, x = i+1, j
	if 0 <= y && y <= m && 0 <= x && x < n && grid[y][x] == 1 {
		res |= 1 << 1
	}
	y, x = i+1, j+1
	if 0 <= y && y <= m && 0 <= x && x < n && grid[y][x] == 1 {
		res |= 1 << 0
	}

	return res
}
