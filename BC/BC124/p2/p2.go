package p2

func lastNonEmptyString(s string) string {
	m := map[byte]int{'a': 0, 'b': 0, 'c': 0, 'd': 0, 'e': 0, 'f': 0, 'g': 0, 'h': 0, 'i': 0, 'j': 0, 'k': 0, 'l': 0, 'm': 0, 'n': 0, 'o': 0, 'p': 0, 'q': 0, 'r': 0, 's': 0, 't': 0, 'u': 0, 'v': 0, 'w': 0, 'x': 0, 'y': 0, 'z': 0}
	maX := 0
	for _, c := range s {
		m[byte(c)]++
		if m[byte(c)] > maX {
			maX = m[byte(c)]
		}
	}
	mm := map[byte]int{}
	for k, v := range m {
		if v == maX {
			mm[k] = v
		}
	}
	n := []byte{}
	for len(mm) > 0 {

		for i := len(s) - 1; i >= 0; i-- {
			if _, ok := mm[byte(s[i])]; ok {
				n = append([]byte{s[i]}, n...)
				delete(mm, byte(s[i]))
				break
			}
		}
	}
	return string(n[:])
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
