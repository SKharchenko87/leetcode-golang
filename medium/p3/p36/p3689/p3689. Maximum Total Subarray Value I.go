package p3689

import "math"

func maxTotalValue(nums []int, k int) int64 {
	maxNums := math.MinInt
	minNums := math.MaxInt
	for _, num := range nums {
		maxNums = max(maxNums, num)
		minNums = min(minNums, num)
	}
	return int64(maxNums-minNums) * int64(k)
}
