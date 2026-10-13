package p3727

import (
	"slices"
)

func maxAlternatingSum(nums []int) int64 {
	n := len(nums)
	slices.SortFunc(nums, func(a, b int) int {
		return abs(b) - abs(a)
	})
	res := int64(0)
	for i := 0; i < n/2; i++ {
		res += int64(nums[i] * nums[i])
		res -= int64(nums[n-i-1] * nums[n-i-1])
	}
	if n%2 == 1 {
		res += int64(nums[n/2] * nums[n/2])
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
