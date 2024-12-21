package p2

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -1 * x
}

const MOD = 1e9 + 7

func countPathsWithXorValue(grid [][]int, k int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][][]int, m)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, 16)
		}
	}

	dp[0][0][grid[0][0]] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for x := 0; x < 16; x++ {
				if i < m-1 {
					y := x ^ grid[i+1][j]
					dp[i+1][j][y] += dp[i][j][x]
					dp[i+1][j][y] %= MOD
				}
				if j < n-1 {
					y := x ^ grid[i][j+1]
					dp[i][j+1][y] += dp[i][j][x]
					dp[i][j+1][y] %= MOD
				}

			}

		}
	}
	return dp[m-1][n-1][k]
}

func countPathsWithXorValue1(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])

	// Создаем 3D DP массив: dp[i][j][xor] - количество путей до ячейки (i, j) с XOR равным `xor`
	dp := make([][][]int, m)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, 1024) // Максимальный XOR для 10-битных чисел
		}
	}

	// Инициализируем стартовую ячейку
	dp[0][0][grid[0][0]] = 1

	// Заполняем DP таблицу
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for xor := 0; xor < 1024; xor++ {
				if dp[i][j][xor] > 0 {
					// Переход вниз
					if i+1 < m {
						nextXor := xor ^ grid[i+1][j]
						dp[i+1][j][nextXor] = (dp[i+1][j][nextXor] + dp[i][j][xor]) % MOD
					}
					// Переход вправо
					if j+1 < n {
						nextXor := xor ^ grid[i][j+1]
						dp[i][j+1][nextXor] = (dp[i][j+1][nextXor] + dp[i][j][xor]) % MOD
					}
				}
			}
		}
	}

	// Возвращаем количество путей с XOR равным k в последней ячейке
	return dp[m-1][n-1][k]
}

func countPathsWithXorValue0(grid [][]int, k int) int {
	m := len(grid)
	n := len(grid[0])
	candidate := [][3]int{{0, 0, grid[0][0]}}
	res := 0
	i := 0
	for len(candidate) != 0 {
		x, y := candidate[i][0], candidate[i][1]
		c := candidate[i][2]
		if x == m-1 && y == n-1 && c == k {
			res++
		}
		if x < m-1 {
			candidate = append(candidate, [3]int{x + 1, y, grid[x+1][y] ^ c})
		}
		if y < n-1 {
			candidate = append(candidate, [3]int{x, y + 1, grid[x][y+1] ^ c})
		}
		candidate = candidate[1:]
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs64(x int64) int64 {
	if x >= 0 {
		return x
	}
	return -1 * x

}

func min64(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func max64(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}
