package p3

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

func isValis(s string) bool {
	al := [26]int{}
	x := 0
	for _, c := range s {
		al[c-'a']++
		x = al[c-'a']
	}
	for _, v := range al {
		if v > 0 && v != x {
			return false
		}
	}
	return true
}

func check(a, b [26]int) bool {
	var x [26]int
	for i := 0; i < 26; i++ {
		x[i] = b[i] - a[i]
	}
	y := 0
	for i := 0; i < 26; i++ {
		if x[i] != 0 {
			if y == 0 {
				y = x[i]
			}
			if y != x[i] {
				return false
			}
		}
	}
	return true
}

func minimumSubstringsInPartition(s string) int {
	l := len(s)
	arr := make([][26]int, l)
	arr[0][s[0]-'a']++
	for i := 1; i < l; i++ {
		for j := 0; j < 26; j++ {
			arr[i][j] = arr[i-1][j]
		}
		arr[i][s[i]-'a']++
	}
	m := map[int][]int{}
	for i := 0; i < l; i++ {
		m[i] = []int{}
		for j := i; j < l; j++ {
			if isValis(s[i : j+1]) {
				m[i] = append(m[i], j)
			}
		}
	}
	var dfs func(start int) int
	dfs = func(start int) int {
		if start == l {
			return 0
		}
		x := 10000
		for i := len(m[start]) - 1; i >= 0; i-- {
			x = min(x, dfs(m[start][i]+1)+1)
		}
		return x
	}
	return dfs(0)
}
