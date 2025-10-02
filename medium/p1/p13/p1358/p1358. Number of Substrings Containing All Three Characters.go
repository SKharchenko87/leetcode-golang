package p1358

func numberOfSubstrings(s string) int {
	var a, b, c int
	count := 0
	l := 0
	for r := 0; r < len(s); r++ {
		switch s[r] {
		case 'a':
			a++
		case 'b':
			b++
		case 'c':
			c++
		}

		for ; a > 0 && b > 0 && c > 0 && l < r; l++ {
			count += len(s) - r
			switch s[l] {
			case 'a':
				a--
			case 'b':
				b--
			case 'c':
				c--
			}
		}
	}
	return count
}
