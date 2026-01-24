package p1877

import (
	"math"
	"sort"
)

func minPairSum(nums []int) int {
	sort.Ints(nums)
	result := math.MinInt
	i := 0
	j := len(nums) - 1
	for i < j {
		result = max(result, nums[i]+nums[j])
		i++
		j--
	}
	return result
}

func minPairSum0(nums []int) int {
	sort.Ints(nums)
	result := math.MinInt
	for i := 0; i < len(nums)/2; i++ {
		result = max(result, nums[i]+nums[len(nums)-i-1])
	}
	return result
}
