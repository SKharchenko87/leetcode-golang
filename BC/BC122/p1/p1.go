package p1

import "sort"

func minimumCost(nums []int) int {
	l := len(nums)
	x := nums[1:l]
	sort.Ints(x)
	return nums[0] + x[0] + x[1]
}
