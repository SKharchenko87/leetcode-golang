package p0611

import (
	"slices"
	"sort"
)

func triangleNumber(nums []int) int {
	l := len(nums)
	if l < 3 {
		return 0
	}
	sort.Ints(nums)
	i := 0
	for ; nums[i] == 0 && i < l-2; i++ {
	}

	count := 0
	for ; i < l-2; i++ {
		k := i + 2
		for j := i + 1; j < l-1; j++ {
			for ; k < l && nums[k] < nums[i]+nums[j]; k++ {
			}
			k--
			count += k - j
		}
	}
	return count
}

func triangleNumber0(nums []int) int {
	l := len(nums)
	if l < 3 {
		return 0
	}
	sort.Ints(nums)
	i := 0
	for ; nums[i] == 0 && i < l-2; i++ {
	}

	count := 0
	for ; i < l-2; i++ {
		for j := i + 1; j < l-1; j++ {
			k, _ := slices.BinarySearch(nums, nums[i]+nums[j])
			count += k - j - 1
		}
	}
	return count
}
