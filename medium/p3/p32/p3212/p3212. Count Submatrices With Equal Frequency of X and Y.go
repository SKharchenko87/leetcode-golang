package p3212

type XY struct {
	x int
	y int
}

func numberOfSubmatrices(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	xy := make([]XY, n)
	var cur XY
	res := 0
	for i := 0; i < n; i++ {
		if grid[0][i] == 'X' {
			cur.x++
		} else if grid[0][i] == 'Y' {
			cur.y++
		}
		xy[i] = cur
		if xy[i].x > 0 && xy[i].x == xy[i].y {
			res++
		}
	}
	for i := 1; i < m; i++ {
		cur.x = 0
		cur.y = 0
		for j := 0; j < n; j++ {
			if grid[i][j] == 'X' {
				cur.x++
			} else if grid[i][j] == 'Y' {
				cur.y++
			}
			xy[j].x += cur.x
			xy[j].y += cur.y
			if xy[j].x > 0 && xy[j].x == xy[j].y {
				res++
			}
		}
	}
	return res
}
