package p3000

func areaOfMaxDiagonal(dimensions [][]int) int {
	maxPreDiag := 0
	maxArea := 0
	for _, dimension := range dimensions {
		currentDiagonal := dimension[0]*dimension[0] + dimension[1]*dimension[1]
		currentArea := dimension[0] * dimension[1]
		if maxPreDiag == currentDiagonal {
			maxArea = max(maxArea, currentArea)
		} else if maxPreDiag < currentDiagonal {
			maxPreDiag = currentDiagonal
			maxArea = currentArea
		}
	}
	return maxArea
}
