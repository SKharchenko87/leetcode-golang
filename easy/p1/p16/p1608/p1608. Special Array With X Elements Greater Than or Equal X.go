package p1608

import "sort"

func specialArray(nums []int) int {
	sort.Ints(nums)
	l := len(nums)
	i := l - 1
	for ; i >= 0; i-- {
		if nums[i] < (l - i) {
			break
		}
	}
	res := l - i - 1
	if res == 0 {
		return -1
	}
	if i > 0 && nums[i] >= res {
		return -1
	}
	return res
}
