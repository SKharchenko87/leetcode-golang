package p2

import "sort"

type Ints interface {
	int | int8 | int16 | int32 | int64
}
type Floats interface {
	float32 | float64
}
type Numeric interface{ Ints | Floats }

func abs[T Numeric](x T) T {
	if x >= 0 {
		return x
	}
	return -1 * x
}

func min[T Numeric](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func max[T Numeric](x, y T) T {
	if x > y {
		return x
	}
	return y
}

func radius(p []int) int64 {
	return int64(p[0])*int64(p[0]) + int64(p[1])*int64(p[1])
}

type normPoint struct {
	mc  int
	tag byte
}

func maxPointsInsideSquare(points [][]int, s string) int {
	l := len(points)
	arr := make([]normPoint, l)
	for i := 0; i < l; i++ {
		arr[i] = normPoint{max(abs(points[i][0]), abs(points[i][1])), s[i]}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].mc < arr[j].mc
	})
	m := map[byte]bool{}
	cnt := 0
	mcPrev := 0
	mcPrevCount := 0
	for _, point := range arr {
		if mcPrev != point.mc {
			mcPrevCount = cnt
		}
		if m[point.tag] {
			return mcPrevCount
		}
		m[point.tag] = true
		cnt++
		mcPrev = point.mc
	}
	return cnt
}
