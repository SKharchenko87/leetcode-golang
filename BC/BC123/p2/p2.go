package p2

import "sort"

func numberOfPairs(points [][]int) int {
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
