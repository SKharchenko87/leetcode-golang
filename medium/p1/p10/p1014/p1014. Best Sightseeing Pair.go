package p1014

import "sort"

func maxScoreSightseeingPair(values []int) int {
	l := len(values)
	maxLeft := values[0]
	res := 0
	for i := 1; i < l; i++ {
		res = max(res, maxLeft+values[i]-i)
		maxLeft = max(maxLeft, values[i]+i)
	}
	return res
}

type Ints interface {
	int | int8 | int16 | int32 | int64
}
type Floats interface {
	float32 | float64
}
type Numeric interface{ Ints | Floats }

func max[T Numeric](x, y T) T {
	if x > y {
		return x
	}
	return y
}

func maxScoreSightseeingPair1(values []int) int {
	l := len(values)
	indexes := make([]int, l)
	for i := range indexes {
		indexes[i] = i
	}
	sort.SliceStable(indexes, func(i, j int) bool {
		return values[indexes[i]] > values[indexes[j]]
	})
	res := 0
	i := 0
	indexI := indexes[i]
	for ; i < l-1 && 2*values[indexes[i]] > res; i++ {
		indexI = indexes[i]
		j := i + 1
		indexJ := indexes[j]
		for ; j < l && values[indexI]+values[indexJ] > res; j++ {
			indexJ = indexes[j]
			if indexI < indexJ {
				if res < values[indexI]+values[indexJ]+indexI-indexJ {
					res = values[indexI] + values[indexJ] + indexI - indexJ
				}
			} else if res < values[indexI]+values[indexJ]+indexJ-indexI {
				res = values[indexI] + values[indexJ] + indexJ - indexI
			}

		}
	}
	return res
}

// TLE
func maxScoreSightseeingPair0(values []int) int {
	l := len(values)
	res := 0
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if res < values[i]+values[j]+i-j {
				res = values[i] + values[j] + i - j
			}
		}
	}
	return res
}
