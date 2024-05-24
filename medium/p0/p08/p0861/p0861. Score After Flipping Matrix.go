package p0861

// ToDo banchmark
func matrixScore(grid [][]int) (sum int) {
	m, n := len(grid), len(grid[0])
	for j := 1; j < n; j++ {
		tmp := 0
		for i := 0; i < m; i++ {
			tmp += 1 - grid[i][0] ^ grid[i][j]
		}
		sum += max(m-tmp, tmp) << (n - 1 - j)
	}
	sum += 1 << (n - 1) * m
	return sum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func matrixScore1(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	arr := make([]int, n)

	for i := 0; i < m; i++ {
		for j := 1; j < n; j++ {
			arr[j] = arr[j] + (1 - grid[i][0] ^ grid[i][j])
		}
		arr[0] += 1
	}

	sum := 1 << (n - 1) * m
	for j := 1; j < n; j++ {
		sum += max(m-arr[j], arr[j]) * (1 << (n - 1 - j))
	}
	return sum
}

func matrixScore0(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	arr := make([]int, m)
	sum := 0
	for i := 0; i < m; i++ {
		tmp := 0
		for j := 0; j < n; j++ {
			tmp += grid[i][j] << (n - 1 - j)
		}
		if grid[i][0] == 0 {
			tmp ^= 1<<n - 1
		}
		arr[i] = tmp
		sum += arr[i]
	}

	for j := 1; j < n; j++ {
		tmpSum := 0
		tmpArr := make([]int, m)
		for i := 0; i < m; i++ {
			tmpArr[i] = arr[i] ^ (1 << (n - 1 - j))
			tmpSum += tmpArr[i]
		}
		if tmpSum > sum {
			sum = tmpSum
			arr = tmpArr
		}
	}
	return sum
}
