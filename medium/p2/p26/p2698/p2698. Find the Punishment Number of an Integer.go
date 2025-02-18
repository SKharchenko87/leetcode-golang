package p2698

func punishmentNumber(n int) int {
	ans := 0
	for i := 1; i <= n; i++ {
		if s := i * i; isMatching(s, i) {
			ans += s
		}
	}
	return ans
}

func isMatching(p, i int) bool {
	var bt func(ts, p, i int) bool
	bt = func(ts, p, i int) bool {
		if p <= 0 {
			if ts == i {
				return true
			}
			return false
		}
		for j := 10; j < 1000000; j = j * 10 {
			if bt(ts+p%j, p/j, i) {
				return true
			}
		}
		return false
	}
	return bt(0, p, i)
}
