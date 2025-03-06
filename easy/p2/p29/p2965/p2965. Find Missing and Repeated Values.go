package p2965

func findMissingAndRepeatedValues(grid [][]int) []int {
	gotSum := 0
	gotSqrSum := 0
	sqrSum := 0
	sum := 0
	i := 1
	for _, row := range grid {
		for _, col := range row {
			gotSqrSum += col * col
			gotSum += col
			sqrSum += i * i
			sum += i
			i++
		}
	}
	sqrDiff := (gotSqrSum - sqrSum) // x^2-y^2
	diff := gotSum - sum            // x-y
	xPlusY := sqrDiff / diff        // x+y
	x := (xPlusY + diff) / 2
	y := (xPlusY - diff) / 2
	return []int{x, y}
}

func findMissingAndRepeatedValues1(grid [][]int) []int {
	n := len(grid)
	m := make(map[int]bool, n-1)
	sum := n * n * (n*n + 1) / 2
	twice := 0
	for _, row := range grid {
		for _, col := range row {
			if m[col] {
				twice = col
			}
			m[col] = true
			sum -= col
		}
	}
	return []int{twice, sum + twice}
}

func findMissingAndRepeatedValues0(grid [][]int) []int {
	n := len(grid)
	m := make(map[int]int, n*n)
	counter := 0
	twice := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			counter = counter + grid[i][j] - (j*n + i + 1)
			m[grid[i][j]]++
			if m[grid[i][j]] == 2 {
				twice = grid[i][j]
			}
		}
	}
	return []int{twice, -1 * (counter - twice)}
}
