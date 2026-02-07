package p3634

import (
	"slices"
	"sort"
)

func minRemoval(nums []int, k int) int {
	n := len(nums)
	sort.Ints(nums)
	startIndex := 0
	endIndex, _ := slices.BinarySearch(nums, nums[startIndex]*k+1)
	res := n - (endIndex - startIndex)
	for startIndex = 1; startIndex < n && endIndex < n; startIndex++ {
		if startIndex == endIndex {
			endIndex++
		}
		for ; endIndex < n && nums[startIndex]*k >= nums[endIndex]; endIndex++ {
		}
		res = min(res, n-(endIndex-startIndex))
	}
	return res
}
