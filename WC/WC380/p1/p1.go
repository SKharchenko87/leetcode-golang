package p1

import (
	"sort"
)

func maxFrequencyElements(nums []int) int {
	l := len(nums)
	m := map[int]int{}
	sort.Ints(nums)
	count := 1
	res := 0
	for i := 1; i < l; i++ {
		if nums[i-1] == nums[i] {
			count++
		} else {
			m[count]++
			if res < count {
				res = count
			}
			count = 1
		}
	}
	m[count]++
	if res < count {
		res = count
	}
	return m[res] * res
}
