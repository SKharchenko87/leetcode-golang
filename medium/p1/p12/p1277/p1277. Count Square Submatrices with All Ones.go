package p1277

func countSquares(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	res := 0
	for j := 0; j < n; j++ {
		res += matrix[0][j]
	}
	for i := 1; i < m; i++ {
		res += matrix[i][0]
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				continue
			}
			matrix[i][j] += min(matrix[i-1][j], matrix[i-1][j-1], matrix[i][j-1])
			res += matrix[i][j]
		}
	}
	return res
}

func countSquares5(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([]int, n+1)
	res := 0
	for i := 1; i <= m; i++ {
		prev := dp[0]
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == 1 {
				prev, dp[j] = dp[j], 1+min(prev, dp[j-1], dp[j])
				res += dp[j]
			} else {
				dp[j] = 0
			}
		}
	}
	return res
}

func countSquares6(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				continue
			}
			if i > 0 && j > 0 {
				matrix[i][j] += min(matrix[i-1][j], matrix[i-1][j-1], matrix[i][j-1])
			}
			res += matrix[i][j]
		}
	}
	return res
}

func countSquares4(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	res := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == 1 {
				dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
				res += dp[i][j]
			}
		}
	}
	return res
}

func countSquares3(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	res := 0
	checkSquare := func(i, j int) int {
		startI, startJ := i, j
		result := 0
	LOOP:
		for ; i < m && j < n; i, j = i+1, j+1 {
			for i0 := startI; i0 <= i; i0++ {
				if matrix[i0][j] == 0 {
					break LOOP
				}
			}
			for j0 := startJ; j0 < j; j0++ {
				if matrix[i][j0] == 0 {
					break LOOP
				}
			}
			result++
		}
		return result
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				res += checkSquare(i, j)
			}
		}
	}
	return res
}

func countSquares2(matrix [][]int) int {
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

//func min(x, y int) int {
//	if x < y {
//		return x
//	}
//	return y
//}

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
