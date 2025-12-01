package p2141

import "math"

func maxRunTime(n int, batteries []int) int64 {
	sum := sums(batteries, math.MaxInt64)
	c := int64(n)
	high := sum / c
	var low, mid int64
	for low <= high {
		mid = low + (high-low)/2
		if sums(batteries, mid) >= mid*c {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low - 1
}

func sums(batteries []int, t int64) int64 {
	var sum int64
	for _, battery := range batteries {
		sum += min(int64(battery), t)
	}
	return sum
}
