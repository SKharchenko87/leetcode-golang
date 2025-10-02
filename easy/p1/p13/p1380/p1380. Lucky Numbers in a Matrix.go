package p1380

import "slices"

// 1 7
// 8 2

func luckyNumbers(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	rowMinMax := 0
	colMaxMin := 100_001
	for i := 0; i < m; i++ {
		x := slices.Min(matrix[i])
		rowMinMax = max(rowMinMax, x)
	}
	for j := 0; j < n; j++ {
		x := matrix[0][j]
		for i := 1; i < m; i++ {
			if x < matrix[i][j] {
				x = matrix[i][j]
			}
		}
		colMaxMin = min(colMaxMin, x)
	}
	if rowMinMax == colMaxMin {
		return []int{rowMinMax}
	}
	return []int{}
}

func luckyNumbers2(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	min0 := matrix[0][0]
	for j := 0; j < n; j++ {
		//min0
		if matrix[0][j] < min0 {
			min0 = matrix[0][j]
		}
	}

	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			//min
			if matrix[i][0] > matrix[i][j] {
				matrix[i][0] = matrix[i][j]
			}
			//max
			if matrix[i][j] > matrix[0][j] {
				matrix[0][j] = matrix[i][j]
			}
		}
	}
	result := make([]int, 1)
	for j := 0; j < n; j++ {
		if min0 == matrix[0][j] {
			result[0] = min0
			return result
		}
	}
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][0] == matrix[0][j] {
				result[0] = matrix[0][j]
				return result
			}
		}
	}
	return []int{}
}

func luckyNumbers1(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	mins := make([]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mins[i] == 0 || matrix[i][j] < mins[i] {
				mins[i] = matrix[i][j]
			}
			if matrix[i][j] > matrix[0][j] {
				matrix[0][j] = matrix[i][j]
			}
		}
	}
	result := make([]int, 1)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mins[i] == matrix[0][j] {
				result[0] = mins[i]
				return result
			}
		}
	}
	return []int{}
}

func luckyNumbers0(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	mins := make([]int, m)
	maxs := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mins[i] == 0 || matrix[i][j] < mins[i] {
				mins[i] = matrix[i][j]
			}
			if matrix[i][j] > maxs[j] {
				maxs[j] = matrix[i][j]
			}
		}
	}
	result := make([]int, 1)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mins[i] == maxs[j] {
				result[0] = mins[i]
				return result
			}
		}
	}
	return []int{}
}
