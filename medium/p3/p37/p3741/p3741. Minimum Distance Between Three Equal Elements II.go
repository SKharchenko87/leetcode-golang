package p3741

import (
	"math"
)

type distance struct {
	distance  int
	lastIndex int
}

func minimumDistance(nums []int) int {
	d := make([]distance, 100_001)
	res := math.MaxInt
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		newLastIndex := i + 1
		if d[n].lastIndex != 0 {
			newDistance := newLastIndex - d[n].lastIndex
			if d[n].distance != 0 {
				res = min(res, (d[n].distance+newDistance)*2)
			}
			d[n].distance = newDistance
		}
		d[n].lastIndex = newLastIndex
	}
	if res == math.MaxInt {
		return -1
	}
	return res
}
