package p1

import "sort"

func areaOfMaxDiagonal(dimensions [][]int) int {
	sort.Slice(dimensions, func(i, j int) bool {
		id := dimensions[i][0]*dimensions[i][0] + dimensions[i][1]*dimensions[i][1]
		jd := dimensions[j][0]*dimensions[j][0] + dimensions[j][1]*dimensions[j][1]
		if id > jd {
			return true
		} else if id == jd {
			return dimensions[i][0]*dimensions[i][1] > dimensions[j][0]*dimensions[j][1]
		} else {
			return false
		}
	})
	return dimensions[0][0] * dimensions[0][1]
}
