package p2779

import (
	"slices"
	"sort"
)

func maximumBeauty(nums []int, k int) int {
	l := len(nums)
	sort.Ints(nums)
	leftIndex := 0
	num := nums[leftIndex]
	rightIndex, _ := slices.BinarySearch(nums, num+k+1)
	rightIndex--
	res := rightIndex - leftIndex + 1
	for ; num <= nums[len(nums)-1]; num++ {
		for ; leftIndex < l-1 && num-k > nums[leftIndex]; leftIndex++ {
		}
		for ; rightIndex < l && num+k >= nums[rightIndex]; rightIndex++ {
		}
		rightIndex--
		if res < rightIndex-leftIndex+1 {
			res = rightIndex - leftIndex + 1
		}
	}
	return res
}

func maximumBeauty0(nums []int, k int) int {
	sort.Ints(nums)
	res := 0
	for num := nums[0]; num <= nums[len(nums)-1]; num++ {
		start, _ := slices.BinarySearch(nums, num-k)
		end, _ := slices.BinarySearch(nums, num+k+1)
		end--
		if res < end-start+1 {
			res = end - start + 1
		}
	}
	return res
}
