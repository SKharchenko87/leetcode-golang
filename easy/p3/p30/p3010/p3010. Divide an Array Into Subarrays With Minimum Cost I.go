package p3010

import "sort"

func minimumCost(nums []int) int {
	b, c := nums[1], nums[2]
	b, c = min(b, c), max(b, c)
	for i := 3; i < len(nums); i++ {
		if nums[i] <= b {
			b, c = nums[i], b
		} else if nums[i] < c {
			c = nums[i]
		}
	}
	return nums[0] + b + c
}

func minimumCost0(nums []int) int {
	sort.Ints(nums[1:])
	return nums[0] + nums[1] + nums[2]
}
