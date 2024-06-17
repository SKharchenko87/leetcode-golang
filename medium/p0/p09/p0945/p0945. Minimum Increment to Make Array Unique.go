package p0945

import "sort"

// 1 1 2 2 3 7
// 1 2 3 4 5 7

// 1 1 1 1 2 2 3 7
// 1 2 3 4 5 7

func minIncrementForUnique(nums []int) int {
	cntAtt := [100001]int{}
	var res int
	for _, num := range nums {
		cntAtt[num]++
	}
	for i := 0; i < 100000; i++ {
		if cntAtt[i] > 1 {
			res += cntAtt[i] - 1
			cntAtt[i+1] += cntAtt[i] - 1
		}
	}
	res += (cntAtt[100000] - 1) * (cntAtt[100000]) / 2
	return res
}

func minIncrementForUnique2(nums []int) int {
	sort.Ints(nums)
	var res int
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			res += nums[i-1] - nums[i] + 1
			nums[i] += nums[i-1] - nums[i] + 1
		}
	}
	return res
}

func minIncrementForUnique1(nums []int) int {
	const LIMIT = 100000
	cntAtt := make(map[int]int, 33000)
	var res int
	for _, num := range nums {
		cntAtt[num]++
	}
	for i := 0; i < LIMIT; i++ {
		if cntAtt[i] > 1 {
			res += cntAtt[i] - 1
			cntAtt[i+1] += cntAtt[i] - 1
		}
	}
	res += (cntAtt[LIMIT] - 1) * (cntAtt[LIMIT]) / 2
	return res
}
