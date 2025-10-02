package p2

import "sort"

func minCost(arr []int, brr []int, k int64) int64 {
	res0 := int64(0)
	for i := 0; i < len(arr); i++ {
		res0 += int64(abs(arr[i] - brr[i]))
	}
	if res0 <= k {
		return res0
	}

	arr2 := make([][2]int, len(arr))
	for i := 0; i < len(arr); i++ {
		arr2[i] = [2]int{i, arr[i]}
	}
	brr2 := make([][2]int, len(brr))
	for i := 0; i < len(brr); i++ {
		brr2[i] = [2]int{i, brr[i]}
	}
	sort.Slice(arr2, func(i, j int) bool {
		return arr2[i][1] < arr2[j][1]
	})
	sort.Slice(brr2, func(i, j int) bool {
		return brr2[i][1] < brr2[j][1]
	})
	flg2 := false
	res := int64(0)
	for i := 0; i < len(arr2); i++ {
		if arr2[i][0] != brr2[i][0] {
			flg2 = true
		}
		res += int64(abs(arr2[i][1] - brr2[i][1]))
	}
	if flg2 {
		res += k
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
