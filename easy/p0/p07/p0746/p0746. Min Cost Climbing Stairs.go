package p0746

import (
	"math"
)

func minCostClimbingStairs(cost []int) int {
	l := len(cost)
	minPath := make([]int, l)
	k := 2
	for i := 0; i < k && i < l; i++ {
		minPath[i] = cost[i]
	}
	for i := k; i < l; i++ {
		prevMin := math.MaxInt
		for j := 1; j <= k && i-j >= 0; j++ {
			if prevMin > minPath[i-j] {
				prevMin = minPath[i-j]
			}
		}
		minPath[i] = cost[i] + prevMin
	}
	res := math.MaxInt
	for i := 1; i <= k; i++ {
		if res > minPath[l-i] {
			res = minPath[l-i]
		}
	}
	return res
}
