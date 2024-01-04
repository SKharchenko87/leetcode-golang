package p0452

import (
	"math"
	"sort"
)

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool { return points[i][0] < points[j][0] })
	lastEnd := math.MaxInt
	count := 1
	for _, v := range points {
		if v[0] <= lastEnd {
			lastEnd = min(lastEnd, v[1])
		} else {
			count++
			lastEnd = v[1]
		}
	}
	return count
}
