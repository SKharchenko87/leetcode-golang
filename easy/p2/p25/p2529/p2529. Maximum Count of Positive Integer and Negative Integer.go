package p2529

import "slices"

func maximumCount(nums []int) int {
	indexNegative, _ := slices.BinarySearch(nums, 0)
	indexPositive, _ := slices.BinarySearch(nums, 1)
	return max(indexNegative, len(nums)-indexPositive)
}
