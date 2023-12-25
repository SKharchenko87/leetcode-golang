package main

import (
	"fmt"
	"sort"
)

func main() {
	points := [][]int{{3, 1}, {9, 0}, {1, 0}, {1, 4}, {5, 3}, {8, 8}}
	fmt.Println(maxWidthOfVerticalArea(points))
}
func maxWidthOfVerticalArea(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	maxArea := 0
	for i := 1; i < len(points); i++ {
		x := points[i][0] - points[i-1][0]
		if maxArea < x {
			maxArea = x
		}
	}
	return maxArea
}
