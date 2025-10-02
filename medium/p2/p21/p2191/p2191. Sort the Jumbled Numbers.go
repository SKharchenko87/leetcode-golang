package p2191

import (
	"sort"
)

type order struct {
	nums        *[]int
	numsOrdered *[]int
}

func (o *order) Len() int {
	return len(*o.nums)
}

func (o *order) Less(i, j int) bool {
	return (*(*o).numsOrdered)[i] < (*(*o).numsOrdered)[j]
}

func (o *order) Swap(i, j int) {
	(*(*o).nums)[i], (*(*o).nums)[j] = (*(*o).nums)[j], (*(*o).nums)[i]
	(*(*o).numsOrdered)[i], (*(*o).numsOrdered)[j] = (*(*o).numsOrdered)[j], (*(*o).numsOrdered)[i]
}

func sortJumbled(mapping []int, nums []int) []int {
	numsOrdered := make([]int, len(nums))
	for i, num := range nums {
		numsOrdered[i] = mapper(num, mapping)
	}
	o := order{&nums, &numsOrdered}
	sort.Stable(&o)
	return nums
}

type order1 struct {
	nums        *[]int
	numsOrdered *[]int
}

func (o order1) Len() int {
	return len(*o.nums)
}

func (o order1) Less(i, j int) bool {
	return (*o.numsOrdered)[i] < (*o.numsOrdered)[j]
}

func (o order1) Swap(i, j int) {
	(*o.nums)[i], (*o.nums)[j] = (*o.nums)[j], (*o.nums)[i]
	(*o.numsOrdered)[i], (*o.numsOrdered)[j] = (*o.numsOrdered)[j], (*o.numsOrdered)[i]
}

func sortJumbledX(mapping []int, nums []int) []int {
	numsOrdered := make([]int, len(nums))
	for i, num := range nums {
		numsOrdered[i] = mapper(num, mapping)
	}
	o := order1{&nums, &numsOrdered}
	sort.Stable(o)
	return nums
}

func sortJumbled0(mapping []int, nums []int) []int {
	sort.SliceStable(nums, func(i, j int) bool {
		return mapper(nums[i], mapping) < mapper(nums[j], mapping)
	})
	return nums
}

func mapper(x int, mapping []int) int {
	if x < 10 {
		return mapping[x]
	}
	res := 0
	pow10 := 1
	for x > 0 {
		res += mapping[x%10] * pow10
		pow10 *= 10
		x /= 10
	}
	return res
}
