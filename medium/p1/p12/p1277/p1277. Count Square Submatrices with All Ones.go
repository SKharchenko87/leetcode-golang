package p1277

func countSquares(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])
	l := min(m, n)
	res := 0
	for i := 0; i < m; i++ {
	LOOP:
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				res++
				var x, y int
				for d := 1; d < l; d++ {
					if j+d >= n || i+d >= m {
						continue LOOP
					}
					x = i + d
					for y = j; y <= j+d; y++ {
						if matrix[x][y] != 1 {
							continue LOOP
						}
					}
					y = j + d
					for x = i; x < i+d; x++ {
						if matrix[x][y] != 1 {
							continue LOOP
						}
					}
					res++
				}

			}
		}
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func countSquares1(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])
	l := min(m, n)
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				res++
				res += c(matrix, i, j, m, n, l)
			}
		}
	}
	return res
}

func c(matrix [][]int, x, y, m, n, l int) int {
	res := 0
	var i, j int
	for d := 1; d < l; d++ {
		if y+d >= n || x+d >= m {
			return res
		}
		i = x + d
		for j = y; j <= y+d; j++ {
			if matrix[i][j] != 1 {
				return res
			}
		}
		j = y + d
		for i = x; i < x+d; i++ {
			if matrix[i][j] != 1 {
				return res
			}
		}
		res++
	}
	return res
}

func countSquares0(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])
	mOne := make(map[int]struct{}, m*n)
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				mOne[i*1000+j] = struct{}{}
			}
		}
	}
	l := min(m, n)

	for i := range mOne {
		res++
	LOOP:
		for k := 1; k < l; k++ {
			for j := i / 1000 * 1000; j < i/1000*1000+k*1000+1000; j += 1000 {
				for a := i % 1000; a < i%1000+k+1; a++ {
					if _, ok := mOne[j+a]; !ok {
						continue LOOP
					}
				}
			}
			res++
		}

	}

	return res
}
