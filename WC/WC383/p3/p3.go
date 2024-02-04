package p3

func resultGrid(image [][]int, threshold int) [][]int {
	m, n := len(image), len(image[0])
	// Initialize
	result, sum, count := make([][]int, m), make([][]int, m), make([][]int, m)
	for i := range result {
		result[i], sum[i], count[i] = make([]int, n), make([]int, n), make([]int, n)
	}

	// Iterate through each potential 3x3 subgrid
	for i := 0; i <= m-3; i++ {
		for j := 0; j <= n-3; j++ {
			if isValidRegion(image, i, j, threshold) {
				avg := calculateAverage(image, i, j)
				// Update sum and count for each pixel in the subgrid
				for di := 0; di < 3; di++ {
					for dj := 0; dj < 3; dj++ {
						sum[i+di][j+dj] += avg
						count[i+di][j+dj]++
					}
				}
			}
		}
	}

	// Calculate the final average for each pixel
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if count[i][j] > 0 {
				result[i][j] = sum[i][j] / count[i][j]
			} else {
				result[i][j] = image[i][j]
			}
		}
	}

	return result
}

func isValidRegion(image [][]int, x, y, threshold int) bool {
	for i := x; i <= x+2; i++ {
		for j := y; j <= y+2; j++ {
			if i > x && abs(image[i][j]-image[i-1][j]) > threshold {
				return false
			}
			if j > y && abs(image[i][j]-image[i][j-1]) > threshold {
				return false
			}
		}
	}
	return true
}

func calculateAverage(image [][]int, x, y int) int {
	sum := 0
	for i := x; i < x+3; i++ {
		for j := y; j < y+3; j++ {
			sum += image[i][j]
		}
	}
	return sum / 9 // Average of a 3x3 region
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -1 * x
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
