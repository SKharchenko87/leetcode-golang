package p3301

import "slices"

func maximumTotalSum(maximumHeight []int) int64 {
	slices.SortFunc(maximumHeight, func(a, b int) int {
		return b - a
	})
	curHeight := maximumHeight[0]
	var sum int64
	sum += int64(curHeight)
	for i := 1; i < len(maximumHeight); i++ {
		if maximumHeight[i] >= curHeight-1 {
			curHeight--
		} else {
			curHeight = maximumHeight[i]
		}
		if curHeight == 0 {
			return -1
		}
		sum += int64(curHeight)
	}
	return sum
}
