package p1

import "strings"

func hasMatch(s string, p string) bool {
	arr := strings.Split(p, "*")
	index0 := strings.Index(s, arr[0])
	index1 := strings.Index(s[index0+len(arr[0]):], arr[1])
	return index1 != -1
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
