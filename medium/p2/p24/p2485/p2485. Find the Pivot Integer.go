package p2485

import "math"

func pivotInteger(n int) int {
	x := math.Sqrt(float64((n + n*n) / 2))
	if x == math.Floor(x) {
		return int(x)
	}
	return -1
}
