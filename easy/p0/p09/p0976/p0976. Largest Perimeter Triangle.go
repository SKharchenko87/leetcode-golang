package p0976

import "sort"

func largestPerimeter(nums []int) int {
	sort.Ints(nums)
	for i := len(nums) - 1; i > 1; i-- {
		p := nums[i-1] + nums[i-2]
		if nums[i] < p {
			return nums[i] + p
		}
	}
	return 0
}
