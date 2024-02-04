package p4

func minimumTimeToInitialState(word string, k int) int {
	l := len(word)
	cnt := 0
	for (cnt+1)*k < l {
		cnt++
		start := (cnt) * k
		subR := word[start:]
		subL := word[:len(subR)]
		if subR == subL {
			return cnt
		}
	}
	cnt++
	return cnt
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -1 * x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs64(x int64) int64 {
	if x >= 0 {
		return x
	}
	return -1 * x

}

func min64(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func max64(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}
