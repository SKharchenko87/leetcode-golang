package p3531

func countCoveredBuildings(n int, buildings [][]int) int {
	rowMins := make([]int, n+1)
	rowMaxs := make([]int, n+1)
	colMins := make([]int, n+1)
	colMaxs := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		rowMins[i] = n + 1
		colMins[i] = n + 1
	}
	for _, building := range buildings {
		row, col := building[0], building[1]
		rowMins[row] = min(rowMins[row], col)
		rowMaxs[row] = max(rowMaxs[row], col)
		colMins[col] = min(colMins[col], row)
		colMaxs[col] = max(colMaxs[col], row)

	}
	result := 0
	for _, building := range buildings {
		row, col := building[0], building[1]
		if rowMins[row] < col && col < rowMaxs[row] {
			if colMins[col] < row && row < colMaxs[col] {
				result += 1
			}
		}
	}
	return result
}

type Line struct {
	minIndex int
	maxIndex int
}

func countCoveredBuildings1(n int, buildings [][]int) int {
	rows := make([]Line, n+1)
	cols := make([]Line, n+1)
	for i := 0; i < n+1; i++ {
		rows[i] = Line{minIndex: n + 1, maxIndex: -1}
		cols[i] = Line{minIndex: n + 1, maxIndex: -1}
	}
	for _, building := range buildings {
		row, col := building[0], building[1]
		rows[row].minIndex = min(rows[row].minIndex, col)
		rows[row].maxIndex = max(rows[row].maxIndex, col)
		cols[col].minIndex = min(cols[col].minIndex, row)
		cols[col].maxIndex = max(cols[col].maxIndex, row)
	}
	result := 0
	for _, building := range buildings {
		row, col := building[0], building[1]
		if rows[row].minIndex < col && col < rows[row].maxIndex &&
			cols[col].minIndex < row && row < cols[col].maxIndex {
			result += 1
		}
	}
	return result
}

func countCoveredBuildings0(n int, buildings [][]int) int {
	rows := make(map[int]Line, n)
	cols := make(map[int]Line, n)
	for _, building := range buildings {
		row, col := building[0], building[1]

		if v, ok := rows[building[0]]; ok {
			rows[row] = Line{min(v.minIndex, col), max(v.maxIndex, col)}
		} else {
			rows[row] = Line{col, col}
		}

		if v, ok := cols[building[1]]; ok {
			cols[col] = Line{min(v.minIndex, row), max(v.maxIndex, row)}
		} else {
			cols[col] = Line{row, row}
		}
	}
	result := 0
	for _, building := range buildings {
		row, col := building[0], building[1]
		if rows[row].minIndex < col && col < rows[row].maxIndex &&
			cols[col].minIndex < row && row < cols[col].maxIndex {
			result += 1
		}
	}
	return result
}
