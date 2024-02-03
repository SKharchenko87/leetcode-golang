package p4

import (
	"fmt"
	"sort"
)

func numberOfPairs(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] < points[j][1]
		}
		return points[i][0] < points[j][0]
	})
	fmt.Println(points)
	return 0
}
