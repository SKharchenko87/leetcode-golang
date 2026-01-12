package p1266

func minTimeToVisitAllPoints(points [][]int) int {
	res := 0
	for i := 1; i < len(points); i++ {
		dx, dy := abs(points[i][0]-points[i-1][0]), abs(points[i][1]-points[i-1][1])
		maxD := max(dx, dy)
		res += maxD
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
