package p1

func possibleStringCount(word string) int {
	cntCurrent := 0
	x := 0
	for i := 1; i < len(word); i++ {
		if word[i] == word[i-1] {
			cntCurrent++
		} else {
			x += cntCurrent
			cntCurrent = 0
		}
	}
	x += cntCurrent
	return x + 1
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
