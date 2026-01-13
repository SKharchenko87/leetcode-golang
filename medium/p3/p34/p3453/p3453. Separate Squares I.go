package p3453

import (
	"math"
	"sort"
)

func separateSquares(squares [][]int) float64 {
	n := len(squares)
	sort.Slice(squares, func(i, j int) bool {
		return squares[i][1] < squares[j][1]
	})

	lowInt, highInt := math.MaxInt, 0
	for i := 0; i < n; i++ {
		lowInt, highInt = min(lowInt, squares[i][1]), max(highInt, squares[i][1]+squares[i][2])
	}
	low := float64(lowInt)
	high := float64(highInt)
	totalArea := getTotalArea(squares, high)
	compareArea := math.Floor(totalArea / 2 * 100000)
	for low < high {
		mid := low + (high-low)/2
		currArea := math.Floor(getTotalArea(squares, mid) * 100000)
		if currArea == compareArea {
			high = mid
		} else if currArea < compareArea {
			low = mid
		} else {
			high = mid
		}
		if high-low < 0.000001 {
			return math.Round(mid*100000) / 100000
		}
	}
	return math.Floor(low*100000) / 100000
}

func getTotalArea(squares [][]int, y float64) float64 {
	var res float64
	for i := 0; i < len(squares) && float64(squares[i][1]) < y; i++ {
		y0 := float64(squares[i][1])
		l := float64(squares[i][2])
		h := min(l, y-y0)
		area := h * l
		res += area
	}
	return res
}
