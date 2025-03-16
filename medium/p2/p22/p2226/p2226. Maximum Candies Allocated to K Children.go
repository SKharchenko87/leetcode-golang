package p2226

import (
	"sort"
)

func maximumCandies(candies []int, k int64) int {
	l := len(candies)
	sort.Ints(candies)
	var left, right int
	right = candies[l-1]
	left = 0
	for left < right {
		mid := left + (right-left)/2
		if check(candies, int(k), mid) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if check(candies, int(k), left) {
		return left
	}

	return left - 1
}

func maximumCandies1(candies []int, k int64) int {
	l := int64(len(candies))
	sort.Ints(candies)
	var left, right int64
	right = int64(candies[l-1])
	left = 0
	for left < right {
		mid := left + (right-left)/2
		if check0(candies, k, mid) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if check0(candies, k, left) {
		return int(left)
	}

	return int(left - 1)
}

func maximumCandies0(candies []int, k int64) int {
	l := int64(len(candies))
	sort.Ints(candies)
	var left, right int64

	left = int64(candies[l-1]) / k
	if l >= k {
		left = max(int64(candies[l-k]), left)
	}
	sum := int64(0)
	for i := max(0, l-k); i < l; i++ {
		sum += int64(candies[i])
	}
	if sum/k == 0 {
		return 0
	}
	right = min(sum/k, int64(candies[l-1]))

	for left < right {
		mid := left + (right-left)/2
		if check0(candies, k, mid) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if check0(candies, k, left) {
		return int(left)
	}

	return int(left - 1)
}

func check0(candies []int, k int64, x int64) bool {
	if x == 0 {
		return true
	}
	y := int64(1)
	for i := len(candies) - 1; i >= 0 && y > 0; i-- {
		y = int64(candies[i]) / x
		k -= y
		if k <= 0 {
			return true
		}
	}
	return false
}

func check(candies []int, k int, x int) bool {
	if x == 0 {
		return true
	}
	y := 1
	for i := len(candies) - 1; i >= 0 && y > 0; i-- {
		y = candies[i] / x
		k -= y
		if k <= 0 {
			return true
		}
	}
	return false
}
