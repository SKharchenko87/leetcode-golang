package p0875

import (
	"math"
)

func minEatingSpeed(piles []int, h int) int {

	l := len(piles)
	maxPile := math.MinInt
	for _, v := range piles {
		if maxPile < v {
			maxPile = v
		}
	}
	if l == h {
		return maxPile
	}
	count := 0

	var check = func(x int) bool {
		count = 0
		for _, v := range piles {
			count += (v + x - 1) / x
			if count > h {
				return false
			}
		}
		return count <= h
	}
	l, r := 0, maxPile
	for r-l > 1 {
		m := (l + r) / 2
		if check(m) {
			r = m
		} else {
			l = m
		}
	}
	return r
}
