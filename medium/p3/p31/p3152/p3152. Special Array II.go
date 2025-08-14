package p3152

import "slices"

func isArraySpecial(nums []int, queries [][]int) []bool {
	l := len(nums)
	pairsFalse := make([]int, 0, l)
	for i := 0; i < l-1; i++ {
		if nums[i]&1 == nums[i+1]&1 {
			pairsFalse = append(pairsFalse, i)
		}
	}
	pairsFalse = append(pairsFalse, 10001)
	lq := len(queries)
	res := make([]bool, lq)
	for i, query := range queries {
		index, _ := slices.BinarySearch(pairsFalse, query[0])
		res[i] = !(pairsFalse[index] < query[1])
	}
	return res
}
