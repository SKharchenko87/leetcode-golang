package p2563

import (
	"slices"
	"sort"
)

func countFairPairs0(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)
	var res int64
	for i, num := range nums {
		indexLower, _ := slices.BinarySearch(nums, lower-num)
		indexUpper, _ := slices.BinarySearch(nums, upper-num+1)
		// println(i, indexLower, indexUpper, indexUpper-indexLower)
		if i > indexLower {
			res += int64(min(i, indexUpper) - indexLower)
		}
	}
	return res
}
