package p2948

import (
	"slices"
	"sort"
)

type MultiSortSlices struct {
	indexes []int
	values  []int
}

func (m *MultiSortSlices) Len() int {
	return len(m.values)
}

func (m *MultiSortSlices) Less(i, j int) bool {
	return m.values[i] < m.values[j]
}

func (m *MultiSortSlices) Swap(i, j int) {
	m.values[i], m.values[j] = m.values[j], m.values[i]
	m.indexes[i], m.indexes[j] = m.indexes[j], m.indexes[i]
}

func lexicographicallySmallestArray(nums []int, limit int) []int {
	n := len(nums)

	indexes := make([]int, n)
	for i := 0; i < n; i++ {
		indexes[i] = i
	}
	sort.Sort(&MultiSortSlices{indexes: indexes, values: nums})
	prevIndex := 0
	for i := 1; i < n; i++ {
		if nums[i]-nums[i-1] > limit {
			slices.Sort(indexes[prevIndex:i])
			prevIndex = i
		}
	}
	slices.Sort(indexes[prevIndex:n])

	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[indexes[i]] = nums[i]
	}

	return res
}

func lexicographicallySmallestArray0(nums []int, limit int) []int {

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
