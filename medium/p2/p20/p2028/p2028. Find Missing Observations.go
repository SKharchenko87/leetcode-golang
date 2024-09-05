package p2028

func missingRolls(rolls []int, mean int, n int) []int {
	left := mean*(n+len(rolls)) - n
	for _, roll := range rolls {
		left -= roll
	}

	if left < 0 || left > 5*n {
		return []int{}
	}

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = min(6, 1+left)
		left -= ans[i] - 1
	}
	return ans
}

func missingRolls3(rolls []int, mean int, n int) []int {
	m := len(rolls)
	x := mean * (m + n)
	for i := 0; i < m; i++ {
		x -= rolls[i]
	}
	y := x / n
	if y < 1 || x > 6*n {
		return []int{}
	}
	res := make([]int, n)
	x %= y * n
	i := 0
	y++
	for ; i < x; i++ {
		res[i] = y
	}
	y--
	for ; i < n; i++ {
		res[i] = y
	}
	return res
}

func missingRolls2(rolls []int, mean int, n int) []int {
	m := len(rolls)
	x := mean * (m + n)
	for i := 0; i < m; i++ {
		x -= rolls[i]
	}
	y := x / n
	if y < 1 || y > 6 || x > 6*n {
		return []int{}
	}
	res := make([]int, n)
	x %= y * n
	i := 0
	for ; i < x; i++ {
		res[i] = y + 1
	}
	for ; i < n; i++ {
		res[i] = y
	}
	return res
}

func missingRolls1(rolls []int, mean int, n int) []int {
	m := len(rolls)
	x := mean * (m + n)
	for i := 0; i < m; i++ {
		x -= rolls[i]
	}
	y := x / n
	if y < 1 || y > 6 || x > 6*n {
		return []int{}
	}
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = y
	}
	x %= y * n
	for i := 0; i < x; i++ {
		res[i]++
	}
	return res
}

func missingRolls0(rolls []int, mean int, n int) []int {
	m := len(rolls)
	sum := 0
	for i := 0; i < m; i++ {
		sum += rolls[i]
	}
	x := mean*(m+n) - sum
	y := x / n
	if y < 1 || y > 6 {
		return []int{}
	}
	res := make([]int, n)
	for i := 0; i < n; i++ {
		o := x % y
		add := y
		if o != 0 {
			if o+y > 6 {
				add = 6
			} else {
				add = o + y
			}
		}
		res[i] = add
		x -= add
	}
	return res
}
