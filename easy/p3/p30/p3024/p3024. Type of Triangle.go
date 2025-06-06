package p3024

import "sort"

func triangleType(nums []int) string {
	sort.Ints(nums)
	if nums[0]+nums[1] <= nums[2] {
		return "none"
	} else if nums[0] == nums[1] && nums[1] == nums[2] {
		return "equilateral"
	} else if nums[0] == nums[1] || nums[0] == nums[2] || nums[2] == nums[1] {
		return "isosceles"
	}
	return "scalene"
}
