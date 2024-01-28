package p2

import (
	"math"
)

func maximumLength(nums []int) int {
	m := map[int]int{}
	for _, v := range nums {
		m[v]++
	}
	count := 1
	for k, v := range m {
		if k == 1 {
			count = max(count, v)
		} else if v >= 2 {
			curCount := 2
			for i := 2; ; i = i * 2 {
				if c, ok := m[int(math.Pow(float64(k), float64(i)))]; ok {
					curCount++
					if c == 1 {
						break
					} else {
						curCount++
					}
				} else {
					break
				}
			}
			count = max(count, curCount)
		}
	}
	if count%2 == 0 {
		count--
	}
	return count
}
