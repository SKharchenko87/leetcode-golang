package p4

import "sort"

func resultArray(nums []int) []int {
	l := len(nums)
	m := make(map[int]int, l)
	for i := 0; i < l; i++ {
		m[nums[i]]++
	}
	keys := make([]int, 0, l)
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for i, key := range keys {
		m[key] = i
	}
	lKeys := len(keys)
	gc1, gc2 := make([]int, lKeys), make([]int, lKeys)

	var greaterCount = func(gc []int, val int) int {
		cnt:=0
		for i := m[val]+1; i<lKeys; i++{
			if
		}
		return lKeys - m[val]
	}

	arr1, arr2 := make([]int, 0, l), make([]int, 0, l)
	arr1, arr2 = append(arr1, nums[0]), append(arr2, nums[1])
	m1, m2 := map[int]int{arr1[0]: 0}, map[int]int{arr2[0]: 0}
	index1, index2 := 0, 0
	i := 2
	for ; i < l; i++ {

	}
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -1 * x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs64(x int64) int64 {
	if x >= 0 {
		return x
	}
	return -1 * x

}

func min64(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func max64(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}
