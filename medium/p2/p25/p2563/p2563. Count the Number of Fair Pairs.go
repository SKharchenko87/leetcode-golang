package p2563

import (
	"slices"
	"sort"
)

func countFairPairs(nums []int, lower int, upper int) int64 {
	l := len(nums)
	sort.Ints(nums)
	var res int64
	indexLower := l - 1
	indexUpper := l - 1
	for i := 0; i < l && i < indexUpper; i++ {
		for ; indexUpper > i && nums[i]+nums[indexUpper] > upper; indexUpper-- {
		}
		res += int64(indexUpper - i)
	}
	for i := 0; i < l && i < indexLower; i++ {
		for ; indexLower > i && nums[i]+nums[indexLower] >= lower; indexLower-- {
		}
		res -= int64(indexLower - i)
	}
	return res
}

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
