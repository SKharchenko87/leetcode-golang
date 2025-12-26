package p3075

import "sort"

func maximumHappinessSum(happiness []int, k int) int64 {
	l := len(happiness)
	sort.Ints(happiness)
	var sum int64
	limit := min(l, k)
	for i := 0; i < limit; i++ {
		if happiness[l-i-1]-i > 0 {
			sum += int64(happiness[l-i-1] - i)
		} else {
			return sum
		}
	}
	return sum
}

func maximumHappinessSum0(happiness []int, k int) int64 {
	n := len(happiness)
	sort.Ints(happiness)
	var res int64
	for i := 0; i < n && i < k; i++ {
		tmp := int64(happiness[n-1-i] - i)
		if tmp > 0 {
			res += tmp
		} else {
			return res
		}
	}
	return res
}
