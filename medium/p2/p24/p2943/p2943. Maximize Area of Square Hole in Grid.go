package p2943

import (
	"sort"
)

func maxSequence(grid []int) int {
	l := len(grid)

	if l == 0 {
		return 0
	} else if l == 1 {
		return 1
	}

	res := 0
	currentSequence := 1
	sort.Ints(grid)
	for i := 1; i < l; i++ {
		if grid[i]-grid[i-1] != 1 {
			res = max(res, currentSequence)
			currentSequence = 0
		}
		currentSequence++
	}
	res = max(res, currentSequence)

	return res
}

func maximizeSquareHoleArea(n int, m int, hBars []int, vBars []int) int {
	hMaxSeq := maxSequence(hBars)
	vMaxSeq := maxSequence(vBars)
	l := min(hMaxSeq, vMaxSeq) + 1
	res := l * l
	return res
}

func maximizeSquareHoleArea0(n int, m int, hBars []int, vBars []int) int {
	hMaxSeq := 0
	currSeq := 1
	if len(hBars) > 0 {
		hMaxSeq = 1
	}

	sort.Ints(hBars)
	for i := 1; i < len(hBars); i++ {
		if hBars[i] == hBars[i-1]+1 {
			currSeq++
		} else {
			hMaxSeq = max(hMaxSeq, currSeq)
			currSeq = 1
		}
	}
	hMaxSeq = max(hMaxSeq, currSeq)

	vMaxSeq := 0
	currSeq = 1
	if len(vBars) > 0 {
		vMaxSeq = 1
	}
	sort.Ints(vBars)
	for i := 1; i < len(vBars); i++ {
		if vBars[i] == vBars[i-1]+1 {
			currSeq++
		} else {
			vMaxSeq = max(vMaxSeq, currSeq)
			currSeq = 1
		}
	}
	vMaxSeq = max(vMaxSeq, currSeq)
	l := min(hMaxSeq, vMaxSeq) + 1
	res := l * l
	return res
}
