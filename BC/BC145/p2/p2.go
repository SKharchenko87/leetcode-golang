package p2

import (
	"slices"
	"sort"
)

func findMinimumTime(strength []int, K int) int {
	//l := len(strength)
	sort.Ints(strength)

	x := 1
	sword := 0
	time := 0
	for len(strength) > 0 {
		index, flg := slices.BinarySearch(strength, sword)
		if flg {
			strength = slices.Delete(strength, index, index+1)
			x += K
			sword = 0
		} else if index > 0 {
			strength = slices.Delete(strength, index-1, index)
			x += K
			sword = 0
		}
		sword += x
		time++
	}

	return time - 1
}

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
