package p2948

import (
	"slices"
	"sort"
)

func lexicographicallySmallestArray(nums []int, limit int) []int {

	l := len(nums)
	clone := slices.Clone(nums)
	sort.Ints(nums)

	m := map[int]int{}
	groups := [][]int{[]int{nums[0]}}
	group := 0
	for i := 1; i < l; i++ {
		if nums[i]-nums[i-1] > limit {
			group++
			groups = append(groups, []int{})
		}
		groups[group] = append(groups[group], nums[i])
		m[nums[i]] = group
	}

	for i := 0; i < l; i++ {
		mp := m[clone[i]]
		nums[i] = groups[mp][0]
		groups[mp] = groups[mp][1:len(groups[mp])]
	}

	return nums

}
