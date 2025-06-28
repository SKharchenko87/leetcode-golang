package p2099

import "sort"

type arr2 struct {
	nums, indexes *[]int
}

func (a arr2) Len() int {
	return len(*a.nums)
}

func (a arr2) Less(i, j int) bool {
	if (*a.nums)[i] == (*a.nums)[j] {
		return (*a.indexes)[i] > (*a.indexes)[j]
	}
	return (*a.nums)[i] < (*a.nums)[j]
}

func (a arr2) Swap(i, j int) {
	(*a.nums)[i], (*a.nums)[j] = (*a.nums)[j], (*a.nums)[i]
	(*a.indexes)[i], (*a.indexes)[j] = (*a.indexes)[j], (*a.indexes)[i]
}

func maxSubsequence(nums []int, k int) []int {
	l := len(nums)
	indexes := make([]int, l)
	for i := 0; i < l; i++ {
		indexes[i] = i
	}
	a := arr2{&nums, &indexes}
	sort.Sort(a)

	res := nums[l-k:]
	indexes = indexes[l-k:]

	b := arr2{&indexes, &res}
	sort.Sort(b)
	return res
}
