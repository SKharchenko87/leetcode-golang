package p0885

type vectorClockWise struct {
	dx, dy int
}

func spiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
	y, x := rStart, cStart
	vectorClockWises := []vectorClockWise{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	l := 0
	result := make([][]int, rows*cols)
	i := 0
	for i < rows*cols {
		for j, dxy := range vectorClockWises {
			if j%2 == 0 {
				l++
			}
			for k := 0; k < l; k++ {
				if 0 <= x && x < cols && 0 <= y && y < rows {
					result[i] = []int{y, x}
					i++
				}
				x += dxy.dx
				y += dxy.dy
			}
		}
	}
	return result
}
