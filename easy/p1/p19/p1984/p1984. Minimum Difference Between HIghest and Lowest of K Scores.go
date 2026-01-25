package p1984

import (
	"math"
	"sort"
)

func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	k--
	res := math.MaxInt
	for i := 0; i < len(nums)-k; i++ {
		res = min(nums[i+k]-nums[i], res)
	}
	return res
}
