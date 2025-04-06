package p0368

import (
	"sort"
)

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	l := len(nums)
	dividers := make([]int, l)
	prevIndexes := make([]int, l)
	maxDivider := -1
	maxDividerIndex := -1
	for i := 0; i < l; i++ {
		//dividers[i]++
		j := i
		if prevIndexes[i] == 0 {
			prevIndexes[i] = i + 1
		}
		for ; j < l && nums[i] == nums[j]; j++ {
			dividers[i]++
			prevIndexes[j] = prevIndexes[i]
		}
		for ; j < l; j++ {
			if nums[j]%nums[i] == 0 {
				if dividers[j] < dividers[i] {
					dividers[j] = dividers[i]
					prevIndexes[j] = i + 1
				}
			}
		}
		if dividers[i] > maxDivider {
			maxDivider = dividers[i]
			maxDividerIndex = i
		}
	}
	index := maxDividerIndex + 1
	res := make([]int, maxDivider)
	for i := maxDivider - 1; i >= 0; i-- {
		res[i] = nums[index-1]
		index = prevIndexes[index-1]
	}
	return res
}

// ToDo зазобраться LIS ( Longest Increasing Subsequence )
func largestDivisibleSubset0(nums []int) []int {
	sort.Ints(nums)
	n := len(nums)
	groupSize := make([]int, n)
	prevElement := make([]int, n)
	maxIndex := 0

	for i := 0; i < n; i++ {
		groupSize[i] = 1
		prevElement[i] = -1
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 && groupSize[j]+1 > groupSize[i] {
				groupSize[i] = groupSize[j] + 1
				prevElement[i] = j
			}
		}
		if groupSize[i] > groupSize[maxIndex] {
			maxIndex = i
		}
	}

	var result []int
	for i := maxIndex; i != -1; i = prevElement[i] {
		result = append([]int{nums[i]}, result...)
	}
	return result
}
