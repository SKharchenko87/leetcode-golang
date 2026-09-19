package p1401

func checkOverlap(radius int, xCenter int, yCenter int, x1 int, y1 int, x2 int, y2 int) bool {
	closestX := clamp(xCenter, x1, x2)
	closestY := clamp(yCenter, y1, y2)

	dx := xCenter - closestX
	dy := yCenter - closestY

	return dx*dx+dy*dy <= radius*radius
}

func clamp(val, min, max int) int {
	if val < min {
		return min
	}
	if val > max {
		return max
	}
	return val
}
