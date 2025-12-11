package p3623

import "sort"

const MOD = 1e9 + 7

func countTrapezoids(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	cntPointOfLine := []int{1}
	index := 0
	for i := 1; i < len(points); i++ {
		if points[i-1][1] == points[i][1] {
			cntPointOfLine[index]++
		} else {
			cntPointOfLine = append(cntPointOfLine, 1)
			index++
		}
	}

	variantTwoPointOfLine := 0

	sum := 0
	res := 0
	for _, pointOfLine := range cntPointOfLine {
		variantTwoPointOfLine = pointOfLine * (pointOfLine - 1) / 2

		if variantTwoPointOfLine == 0 {
			continue
		}

		res = (res + sum*variantTwoPointOfLine) % MOD
		sum += variantTwoPointOfLine
	}
	return res
}

// TLE
func countTrapezoids0(points [][]int) int {
	cntPointOfLine := map[int]int{}
	for _, point := range points {
		cntPointOfLine[point[1]]++
	}
	l := len(cntPointOfLine)
	variantTwoPointOfLine := make([]int, l)
	index := 0
	for _, pointOfLine := range cntPointOfLine {
		variantTwoPointOfLine[index] = pointOfLine * (pointOfLine - 1) / 2
		index++
	}
	res := 0
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			res = (res + variantTwoPointOfLine[i]*variantTwoPointOfLine[j]) % MOD
		}
	}
	return res
}
