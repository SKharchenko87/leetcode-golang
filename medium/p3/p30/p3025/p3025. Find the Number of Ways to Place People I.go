package p3025

import "sort"

func numberOfPairs(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] > points[j][1]
		}
		return points[i][0] < points[j][0]
	})
	result := 0
	l := len(points)
	for i := 0; i < l-1; i++ {
	LOOP:
		for j := i + 1; j < l; j++ {
			if points[i][1] >= points[j][1] {
				for k := i + 1; k < j; k++ {
					if points[i][1] >= points[k][1] && points[k][1] >= points[j][1] {
						continue LOOP
					}
				}
				result++
			}

		}
	}
	return result
}

type point struct {
	x, y int
}

func newPoint(arr []int) point {
	return point{arr[0], arr[1]}
}

func numberOfPairs1(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] > points[j][1]
		}
		return points[i][0] < points[j][0]
	})
	result := 0
	l := len(points)
	for i := 0; i < l-1; i++ {
		A := newPoint(points[i])
		for j := i + 1; j < l; j++ {
			B := newPoint(points[j])
			if A.y >= B.y {
				result++
				for k := i + 1; k < j; k++ {
					P := newPoint(points[k])
					if A.y >= P.y && P.y >= B.y {
						result--
						break
					}
				}
			}

		}
	}

	return result
}

/*


























 */

func numberOfPairs0(points [][]int) int {
	l := len(points)
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] > points[j][1]
		}
		return points[i][0] < points[j][0]
	})
	cnt := 0
	for i := 0; i < l-1; i++ {
		ChisatoY := points[i][1]
		for j := i + 1; j < l; j++ {
			TakinaY := points[j][1]
			if ChisatoY < TakinaY {
				continue
			}
			flg := 1
			for k := i + 1; k < j; k++ {
				if TakinaY <= points[k][1] && points[k][1] <= ChisatoY {
					flg = 0
					break
				}
			}
			cnt += flg
		}
	}
	return cnt
}
