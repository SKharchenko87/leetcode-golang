package p1593

func maxUniqueSplit(s string) int {
	l := len(s)
	m := make(map[string]bool, l)
	var backtracking func(start int, lengthSubString int) int
	backtracking = func(start int, lengthSubString int) int {
		if m[s[start:start+lengthSubString]] {
			return 0
		}
		m[s[start:start+lengthSubString]] = true

		maxs := 0
		for i := 1; start+lengthSubString+i <= l; i++ {
			x := backtracking(start+lengthSubString, i)
			maxs = max(maxs, x)
		}
		delete(m, s[start:start+lengthSubString])
		return 1 + maxs

	}
	maxs := 0
	for i := 1; i <= l; i++ {
		x := backtracking(0, i)
		maxs = max(maxs, x)
	}
	return maxs
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxUniqueSplit0(s string) int {
	m := map[string]int{}
	l := len(s)
	var backtracking func(start int, lengthSubString int) int
	backtracking = func(start int, lengthSubString int) int {
		if m[s[start:start+lengthSubString]] > 0 {
			return 0
		}
		m[s[start:start+lengthSubString]]++

		maxs := 0
		for i := 1; start+lengthSubString+i <= l; i++ {
			x := backtracking(start+lengthSubString, i)
			maxs = max(maxs, x)
		}
		m[s[start:start+lengthSubString]]--
		return 1 + maxs

	}
	maxs := 0
	for i := 1; i <= l; i++ {
		x := backtracking(0, i)
		maxs = max(maxs, x)
	}
	return maxs
}
