package p3020

import (
	"sort"
)

func maximumLength(nums []int) int {
	n := len(nums)
	m := make(map[int]int, n)
	number := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if _, ok := m[nums[i]]; !ok {
			number = append(number, nums[i])
		}
		m[nums[i]]++
	}
	sort.Ints(number)
	cnt := 1
	i := 0
	if number[0] == 1 {
		cnt = m[1]
		if cnt%2 == 0 {
			cnt--
		}
		i++
	}
	for ; i < len(number); i++ {
		num := number[i]
		cur := 0
		for m[num] > 1 {
			cur++
			num *= num
		}
		if m[num] == 1 {
			cur = cur*2 + 1
		} else {
			cur = cur*2 - 1
		}
		cnt = max(cnt, cur)
	}
	return cnt
}
