package p3

import "sort"

func minimumArrayLength(nums []int) int {
	l := len(nums)
	sort.Ints(nums)
	x := nums[0]
	if x == 1 {
		count := 0
		for i := 0; i < l; i++ {
			if x == nums[i] {
				count++
			} else {
				break
			}
		}
		count++
		return count / 2
	} else {

	}

}
