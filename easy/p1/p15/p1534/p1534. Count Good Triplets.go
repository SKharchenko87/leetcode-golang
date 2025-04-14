package p1534

func countGoodTriplets(arr []int, a int, b int, c int) int {
	l := len(arr)
	count := 0
	for i := 0; i < l-2; i++ {
		for j := i + 1; j < l-1; j++ {
			if abs(arr[i]-arr[j]) <= a {
				for k := j + 1; k < l; k++ {
					if abs(arr[j]-arr[k]) <= b && abs(arr[i]-arr[k]) <= c {
						count++
					}
				}
			}
		}
	}
	return count
}

func countGoodTriplets0(arr []int, a int, b int, c int) int {
	l := len(arr)
	m := make([][]int, l)
	for i := 0; i < l; i++ {
		m[i] = make([]int, l)
		for j := i + 1; j < l; j++ {
			m[i][j] = abs(arr[i] - arr[j])
		}
	}
	count := 0
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			for k := j + 1; k < l; k++ {
				if m[i][j] <= a && m[j][k] <= b && m[i][k] <= c {
					count++
				}
			}
		}
	}
	return count
}

type Ints interface {
	int | int8 | int16 | int32 | int64
}
type Floats interface {
	float32 | float64
}
type Numeric interface{ Ints | Floats }

func abs[T Numeric](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
