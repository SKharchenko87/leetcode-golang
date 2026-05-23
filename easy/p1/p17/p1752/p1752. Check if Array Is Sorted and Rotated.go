package p1752

import "slices"

func check(nums []int) bool {
	cnt := 0
	n := len(nums)

	for i := 0; i < n && cnt < 2; i++ {
		if nums[i] > nums[(i+1)%n] {
			cnt++
		}
	}

	return cnt < 2
}

func check1(nums []int) bool {
	cnt := 1
	_ = slices.IsSortedFunc(nums, func(a, b int) int {
		if a < b {
			cnt--
		}
		return cnt
	})
	return cnt == 1 || cnt == 0 && nums[0] >= nums[len(nums)-1]
}

func check0(nums []int) bool {
	l := len(nums)
	cntSplit := 0
	if nums[0] >= nums[l-1] {
		cntSplit = 1
	}
	for i := 0; i < l-1; i++ {
		if nums[i] > nums[i+1] {
			cntSplit--
			if cntSplit < 0 {
				return false
			}
		}
	}
	return true
}
