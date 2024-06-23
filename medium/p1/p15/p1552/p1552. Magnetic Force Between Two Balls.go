package p1552

import "sort"

// 1 2 3 4 5 6 7 8
// 3 - 3
// 4 - 2
// 5 - 1

// 1 2 3 4 5 6 7 8 9
// 5 - 2

func maxDistance(position []int, m int) int {
	n := len(position)
	sort.Ints(position)
	start := position[0]
	end := position[n-1]
	if m == 2 {
		return end - start
	}
	left := 1
	right := (end-start)/(m-1) + 1
	res := 1
loop:
	for left < right {
		mid := (left + right) / 2
		cnt := 0
		prevBasket := start - mid
		for j := 0; j < n; j++ {
			if position[j] >= prevBasket+mid {
				cnt++
				prevBasket = position[j]
				if cnt >= m {
					left = mid + 1
					if mid > res {
						res = mid
					}
					continue loop
				}
			}
		}
		right = mid
	}
	return res
}
