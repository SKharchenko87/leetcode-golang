package p1509

import (
	"math"
	"sort"
)

func minDifference(nums []int) int {
	l := len(nums)
	if l <= 4 {
		return 0
	}
	min0, min1, min2, min3 := math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt
	max0, max1, max2, max3 := math.MinInt, math.MinInt, math.MinInt, math.MinInt
	for i := 0; i < l; i++ {
		if !(nums[i] >= min3) {
			if nums[i] <= min0 {
				min0, min1, min2, min3 = nums[i], min0, min1, min2
			} else if nums[i] <= min1 {
				min1, min2, min3 = nums[i], min1, min2
			} else if nums[i] <= min2 {
				min2, min3 = nums[i], min2
			} else {
				min3 = nums[i]
			}
		}

		if !(nums[i] <= max0) {
			if nums[i] >= max3 {
				max0, max1, max2, max3 = max1, max2, max3, nums[i]
			} else if nums[i] >= max2 {
				max0, max1, max2 = max1, max2, nums[i]
			} else if nums[i] >= max1 {
				max0, max1 = max1, nums[i]
			} else {
				max0 = nums[i]
			}
		}

	}
	return min(min(max0-min0, max1-min1), min(max2-min2, max3-min3))
}

func minDifference0(nums []int) int {
	l := len(nums)
	if l <= 4 {
		return 0
	}
	sort.Ints(nums)
	minValue := math.MaxInt
	for i := 0; i < 4; i++ {
		minValue = min(minValue, nums[i+l-4]-nums[i])
	}
	return minValue
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
