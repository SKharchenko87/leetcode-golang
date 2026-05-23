package p2833

func furthestDistanceFromOrigin(moves string) int {
	l, r := 0, 0
	for _, ch := range moves {
		switch ch {
		case 'L':
			l++
		case 'R':
			r++
		}
	}
	return len(moves) - min(l, r)*2
}

func furthestDistanceFromOrigin0(moves string) int {
	l, r, u := 0, 0, 0
	for _, ch := range moves {
		switch ch {
		case '_':
			u++
		case 'L':
			l++
		case 'R':
			r++
		}
	}
	return abs(l-r) + u
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
