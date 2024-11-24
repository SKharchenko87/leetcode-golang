package p1957

import "math"

func maxMatrixSum(matrix [][]int) int64 {
	var result int64
	var cntNegative int
	minAbsVal := math.MaxInt
	for _, row := range matrix {
		for _, val := range row {
			if val < 0 {
				cntNegative++
				result -= int64(val)
				minAbsVal = min(minAbsVal, -val)
			} else if val > 0 {
				result += int64(val)
				minAbsVal = min(minAbsVal, val)
			} else { // val == 0
				minAbsVal = 0
			}

		}
	}
	if cntNegative%2 == 1 {
		result += -int64(minAbsVal) * 2
	}
	return result
}
