package p3068

import "math"

func maximumValueSum(nums []int, k int, edges [][]int) int64 {
	var res int64
	changes := 0
	minDiff := math.MaxInt32

	for _, num := range nums {
		xorNum := num ^ k

		if num < xorNum {
			changes++
			res += int64(xorNum)
		} else {
			res += int64(num)
		}

		diff := abs(num - xorNum)
		if diff < minDiff {
			minDiff = diff
		}
	}

	if changes%2 == 1 {
		return res - int64(minDiff)
	} else {
		return res
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
