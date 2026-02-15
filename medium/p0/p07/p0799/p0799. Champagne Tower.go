package p0799

func champagneTower(poured int, query_row int, query_glass int) float64 {
	pouredUpGlass := [100]float64{}
	pouredDownGlass := [100]float64{}
	pouredUpGlass[0] = float64(poured)
	for row := 1; row <= query_row; row++ {
		pouredDownGlass[0] = 0
		for glass := 0; glass < row; glass++ {
			pouredDownGlass[glass+1] = 0
			if pouredUpGlass[glass] > 1 {
				x := (pouredUpGlass[glass] - 1) / 2
				pouredDownGlass[glass] += x
				pouredDownGlass[glass+1] = x
			}
		}
		pouredUpGlass, pouredDownGlass = pouredDownGlass, pouredUpGlass
	}
	return min(pouredUpGlass[query_glass], 1)
}
