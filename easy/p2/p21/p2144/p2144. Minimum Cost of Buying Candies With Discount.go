package p2144

import (
	"slices"
	"sort"
)

func minimumCost(cost []int) int {
	sort.Ints(cost)
	result := 0
	i := 0
	for _, c := range slices.Backward(cost) {
		if i%3 != 2 {
			result += c
		}
		i++
	}
	return result
}

func minimumCost0(cost []int) int {
	sort.Ints(cost)
	result := 0
	i := 0
	for _, c := range slices.Backward(cost) {
		if i%3 != 2 {
			result += c
		}
		i++
	}
	return result
}
