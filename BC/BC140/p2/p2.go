package p2

import "slices"

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -1 * x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs64(x int64) int64 {
	if x >= 0 {
		return x
	}
	return -1 * x

}

func min64(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func max64(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

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
