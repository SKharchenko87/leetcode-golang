package p3075

import "sort"

func maximumHappinessSum(happiness []int, k int) int64 {
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
