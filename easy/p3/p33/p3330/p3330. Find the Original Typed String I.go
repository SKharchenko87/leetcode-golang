package p3330

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

func possibleStringCount0(word string) int {
	prev := ' '
	count := 1
	res := 1
	for _, ch := range word {
		if ch == prev {
			count++
		} else {
			prev = ch
			res += count - 1
			count = 1
		}
	}
	res += count - 1
	return res
}
