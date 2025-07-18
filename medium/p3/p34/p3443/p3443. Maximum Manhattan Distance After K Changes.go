package p3443

func maxDistance(s string, k int) int {
	v, h := 0, 0
	res := 0
	y := 2 * k
	for i := 0; i < len(s); i++ {
		if s[i] == 'N' {
			v++
		} else if s[i] == 'S' {
			v--
		} else if s[i] == 'W' {
			h--
		} else {
			h++
		}
		x := abs(h) + abs(v)

		res = max(res, min(x+y, i+1))
	}
	return res
}

func maxDistance1(s string, k int) int {
	res := 0
	for m, n, i := 0, 0, 0; i < len(s); i++ {
		if s[i] == 'N' {
			m++
		} else if s[i] == 'S' {
			m--
		} else if s[i] == 'W' {
			n++
		} else {
			n--
		}
		a, b := m, n
		if m < 0 {
			a = -m
		}
		if n < 0 {
			b = -n
		}
		res = max(res, min(a+b+k+k, i+1))
	}
	return res
}

var m = map[int32][2]int{'N': {0, 1}, 'S': {0, -1}, 'E': {1, 0}, 'W': {-1, 0}}

func maxDistance0(s string, k int) int {
	current := [2]int{0, 0}
	res := 0
	for i, ch := range s {
		current = [2]int{current[0] + m[ch][0], current[1] + m[ch][1]}
		x := abs(current[0]) + abs(current[1])

		if i-x > 2*k {
			x += 2 * k
		} else {
			x = i + 1

		}
		res = max(res, x)
	}
	return res
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -1 * x
}
