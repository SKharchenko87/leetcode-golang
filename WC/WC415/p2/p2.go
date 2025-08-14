package p2

import (
	"math"
	"slices"
)

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -1 * x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs64(x int64) int64 {
	if x >= 0 {
		return x
	}
	return -1 * x

}

func min64(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func max64(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

type Element struct {
	val   int
	index int
}

// ToDo
func maxScore(a []int, b []int) int64 {
	l := len(b)
	length := min(16, l)
	top := make([]Element, length)
	for i := 0; i < length; i++ {
		top[i].val = b[i]
		top[i].index = i
	}
	slices.SortFunc(top, func(a, b Element) int {
		return a.val - b.val
	})
	for i := length + 1; i < l; i++ {
		if b[i] > top[0].val {
			indexInsert, _ := slices.BinarySearchFunc(top, b[i], func(element Element, i int) int {
				return element.val - i
			})
			top = slices.Insert(top, indexInsert, Element{b[i], i})
			top = top[1:]
		}
	}
	slices.SortFunc(top, func(a, b Element) int {
		return a.index - b.index
	})

	maxSum := int64(math.MinInt64)
	for i0 := 0; i0 < length-3; i0++ {
		tmp := int64(a[0] * top[i0].val)
		for i1 := i0 + 1; i1 < length-2; i1++ {
			tmp += int64(a[1] * top[i1].val)
			for i2 := i1 + 1; i2 < length-1; i2++ {
				tmp += int64(a[2] * top[i2].val)
				for i3 := i2 + 1; i3 < length-0; i3++ {
					tmp += int64(a[3] * top[i3].val)
					if maxSum < tmp {
						maxSum = tmp
					}
					tmp -= int64(a[3] * top[i3].val)
				}
				tmp -= int64(a[2] * top[i2].val)
			}
			tmp -= int64(a[1] * top[i1].val)
		}
	}
	return maxSum
}
