package p2

func countSubmatrices(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	sum := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sum += grid[i][j]
		}
	}
	cnt := 0
	for a := n - 1; a >= 0; a-- {
		tmp_sum := 0
		for i := m; i >= 1; i-- {
			if sum-tmp_sum <= k {
				//fmt.Println(a, i, sum, tmp_sum)
				cnt += i
				break
			}
			for j := 0; j <= a; j++ {
				tmp_sum += grid[i-1][j]
			}
		}
		//println(cnt, sum)
		for i := m - 1; i >= 0; i-- {
			sum -= grid[i][a]
		}
	}
	return cnt
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
