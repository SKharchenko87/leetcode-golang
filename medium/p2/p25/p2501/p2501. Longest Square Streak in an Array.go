package p2501

import (
	"math"
	"slices"
	"sort"
)

func longestSquareStreak(nums []int) int {
	l := len(nums)

	var res int
	m := make(map[int]struct{}, l)
	for _, num := range nums {
		m[num] = struct{}{}
	}
	for k := range m {
		r := 1
		y := k
		x := int(math.Sqrt(float64(y)))
		for x*x == y {
			if _, ok := m[x]; ok {
				r++
				y = x
				delete(m, x)
			} else {
				break
			}
			x = int(math.Sqrt(float64(y)))
		}
		i := k * k
		for {
			if _, ok := m[i]; ok {
				r++
				delete(m, i)
				i *= i
			} else {
				break
			}
		}
		delete(m, k)
		if res < r {
			res = r
		}
	}
	if res < 2 {
		return -1
	}
	return res
}

func longestSquareStreak0(nums []int) int {
	sort.Ints(nums)
	var res int
	for i := 0; i < len(nums)-1; i++ {
		x := nums[i] * nums[i]
		r := 1
		_, flg := slices.BinarySearch(nums, x)
		for flg {
			r++
			x *= x
			_, flg = slices.BinarySearch(nums, x)
		}
		if res < r {
			res = r
		}
	}
	if res < 2 {
		return -1
	}
	return res
}
