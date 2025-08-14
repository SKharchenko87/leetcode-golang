package p1460

import (
	"slices"
	"sort"
)

func canBeEqual(target []int, arr []int) bool {
	digits := make([]int, 1001)
	for _, t := range target {
		digits[t]++
	}
	for _, t := range arr {
		digits[t]--
		if digits[t] < 0 {
			return false
		}
	}
	return true
}

func canBeEqual1(target []int, arr []int) bool {
	digits := make([]int, 1000)
	for _, t := range target {
		digits[t-1]++
	}
	for _, t := range arr {
		digits[t-1]--
		if digits[t-1] < 0 {
			return false
		}
	}
	return true
}

func canBeEqual0(target []int, arr []int) bool {
	sort.Ints(target)
	sort.Ints(arr)
	return slices.Compare(target, arr) == 0
}
