package p1653

import (
	"slices"
)

func minimumDeletions(s string) int {
	l := len(s)
	if l < 2 {
		return 0
	}
	costs := make([]int, l+1)
	aCnt, bCnt := 0, 0
	for i := 0; i < l; i++ {
		if s[i] != 'a' {
			aCnt++
		}
		costs[i+1] += aCnt
		if s[l-1-i] != 'b' {
			bCnt++
		}
		costs[l-1-i] += bCnt
		//fmt.Println(costs)
	}
	return slices.Min(costs)
}
