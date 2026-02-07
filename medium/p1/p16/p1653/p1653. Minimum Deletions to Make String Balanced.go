package p1653

import (
	"slices"
)

func minimumDeletions(s string) int {
	l := len(s)
	cntA := 0
	i := 0
	for ; i < l; i++ {
		if s[i] == 'a' {
			cntA++
		}
	}
	if cntA == l || cntA == 0 {
		return 0
	}

	i--
	cntB := 0
	maxLen := cntA
	for ; i >= 0; i-- {
		if s[i] == 'b' {
			cntB++
		} else {
			cntA--
		}
		maxLen = max(maxLen, cntA+cntB)
	}
	return l - maxLen
}

func minimumDeletions2(s string) int {
	l := len(s)
	cnts := make([]int, l)
	prev := 0
	i := 0
	for ; i < l; i++ {
		if s[i] == 'a' {
			prev++
		}
		cnts[i] = prev
	}
	i--
	prev = 0
	for ; i >= 0; i-- {
		if s[i] == 'b' {
			prev++
		}
		cnts[i] += prev
	}
	return l - slices.Max(cnts)
}

func minimumDeletions1(s string) int {
	l := len(s)
	cnts := make([]int, l)
	prev := 0
	i := 0
	for ; i < l; i++ {
		if s[i] == 'a' {
			prev++
		}
		cnts[i] = prev
	}
	i--
	prev = 0
	maxLen := 0
	for ; i >= 0; i-- {
		if s[i] == 'b' {
			prev++
		}
		cnts[i] += prev
		maxLen = max(maxLen, cnts[i])
	}
	return l - maxLen
}

/*
*
*
*
*
*
*
*
*
*
*
*
*
*
*
*
*
*
*
*
*
*
*
*
*
*
*
 */
func minimumDeletions0(s string) int {
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
