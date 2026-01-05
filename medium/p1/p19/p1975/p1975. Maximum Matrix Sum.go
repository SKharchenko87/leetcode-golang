package p1975

import "math"

func maxMatrixSum(matrix [][]int) int64 {
	n := len(matrix)
	var res int64
	var minValue int64 = math.MaxInt64
	haveZero := false
	haveNegative := false
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			x := abs(matrix[i][j])
			minValue = min(minValue, x)
			res += x
			haveZero = haveZero || matrix[i][j] == 0
			if matrix[i][j] < 0 {
				haveNegative = !haveNegative
			}
		}
	}
	if haveNegative && !haveZero {
		res -= minValue * 2
	}
	return res
}

func maxMatrixSum0(matrix [][]int) int64 {
	n := len(matrix)
	var res int64
	var countNegativeDigits int64
	var minValue int64 = math.MaxInt64
	haveZero := false
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			x := abs(matrix[i][j])
			minValue = min(minValue, x)
			res += x
			if matrix[i][j] == 0 {
				haveZero = true
			} else if matrix[i][j] < 0 {
				countNegativeDigits++
			}
		}
	}
	if countNegativeDigits&1 == 1 && !haveZero {
		res -= minValue * 2
	}
	return res
}

func abs(x int) int64 {
	if x < 0 {
		x = -x
	}
	return int64(x)
}
