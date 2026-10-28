package p3882

import "math/bits"

// Решение через биты
func minCost(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][16]uint64, n)
	//wordIdx := grid[0][0] >> 6
	//bitIndex := grid[0][0] & 0b0011_1111
	wordIdx, bitIndex := get(grid[0][0])
	dp[0][wordIdx] |= uint64(1) << bitIndex
	for j := 1; j < n; j++ {
		setCost(dp[j-1], &dp[j], grid[0][j])
	}

	dpCurr := make([][16]uint64, n)
	for i := 1; i < m; i++ {
		clear(dpCurr)
		// вычисляем возможные цены для первого элемента в строке
		setCost(dp[0], &dpCurr[0], grid[i][0])

		// проходим по всей строке
		for j := 1; j < n; j++ {
			var combined [16]uint64
			for w := 0; w < 16; w++ {
				combined[w] = dp[j][w] | dpCurr[j-1][w] //  пришли сверху или слева
			}
			setCost(combined, &dpCurr[j], grid[i][j])
		}

		// текущая строка цен становиться предыдущей
		dp, dpCurr = dpCurr, dp
	}
	// ищем минимальное значение
	for wordIdx = 0; wordIdx < 16; wordIdx++ {
		if dp[n-1][wordIdx] > 0 {
			bitIndex = bits.TrailingZeros64(dp[n-1][wordIdx])
			return set(wordIdx, bitIndex)
		}
	}
	return 0
}

func setCost(prev [16]uint64, res *[16]uint64, n int) {
	for wordIdx := 0; wordIdx < 16; wordIdx++ {
		for prev[wordIdx] > 0 {
			bitIndex := bits.TrailingZeros64(prev[wordIdx])
			prev[wordIdx] ^= uint64(1) << bitIndex
			x := set(wordIdx, bitIndex) ^ n
			currWordIdx, currBitIndex := get(x)
			res[currWordIdx] = res[currWordIdx] | uint64(1)<<currBitIndex
		}
	}
}

func get(x int) (wordIdx, bitIndex int) {
	return x >> 6, x & 0b0011_1111
}

func set(wordIdx, bitIndex int) int {
	return wordIdx<<6 + bitIndex
}

func minCost0(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][1024]bool, n)
	dp[0][grid[0][0]] = true // цена в начальной позиции равна значению
	// заполняем возможные цены для первой строки
	for j := 1; j < n; j++ {
		for k := 0; k < 1024; k++ {
			if dp[j-1][k] {
				dp[j][k^grid[0][j]] = true
			}
		}
	}
	dpCurr := make([][1024]bool, n) // текущие возможные цены
	for i := 1; i < m; i++ {
		clear(dpCurr) // очищаем для каждой строки
		// вычисляем возможные цены для первого элемента в строке
		for k := 0; k < 1024; k++ {
			if dp[0][k] {
				dpCurr[0][k^grid[i][0]] = true
			}
		}
		// проходим по всей строке
		for j := 1; j < n; j++ {
			for k := 0; k < 1024; k++ {
				// проверяем верхнюю цену и цену слева
				if dp[j][k] || dpCurr[j-1][k] {
					dpCurr[j][k^grid[i][j]] = true
				}
			}
		}
		// текущая строка цен становиться предыдущей
		dp, dpCurr = dpCurr, dp
	}

	// ищем минимальное значение
	for k := 0; k < 1024; k++ {
		if dp[n-1][k] {
			return k
		}
	}
	return 0
}
