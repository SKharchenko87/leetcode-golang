package p0594

import "sort"

func findLHS(nums []int) int {
	l := len(nums)
	res := 0
	sort.Ints(nums)
	cntPrev := 0
	prev := nums[0]
	i := 0
	for ; i < l && nums[i] == prev; i++ {
		cntPrev++
	}
	if i == l {
		return 0
	}
	cntCurr := 0
	curr := nums[i]
	for ; i < l; i++ {
		if nums[i] == curr {
			cntCurr++
		} else {
			if curr-prev == 1 {
				res = max(res, cntCurr+cntPrev)
			}
			cntPrev = cntCurr
			prev = curr
			curr = nums[i]
			cntCurr = 1
		}
	}
	if curr-prev == 1 {
		res = max(res, cntCurr+cntPrev)
	}
	return res
}

func findLHS0(nums []int) int {
	m := make(map[int]int, len(nums))
	for _, v := range nums {
		m[v]++
	}
	res := 0
	for k, v := range m {
		if val, ok := m[k+1]; ok {
			res = max(res, v+val)
		}
	}
	return res
}
