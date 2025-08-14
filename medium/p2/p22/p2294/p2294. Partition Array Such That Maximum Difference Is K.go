package p2294

import (
	"slices"
	"sort"
)

func partitionArray(nums []int, k int) int {
	sort.Ints(nums)
	res := 0
	minVal := nums[0]
	i := 0
	for {
		j, _ := slices.BinarySearch(nums[i:], minVal+k+1)
		res++
		i += j
		if i < len(nums) {
			minVal = nums[i]
		} else {
			break
		}
	}
	return res
}

func partitionArray0(nums []int, k int) int {
	sort.Ints(nums)
	res := 0
	minVal := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]-minVal > k {
			res++
			minVal = nums[i]
		}
	}
	return res + 1
}
