package p1358

func numberOfSubstrings(s string) int {
	lastIndexA, lastIndexB, lastIndexC := -1, -1, -1
	prev := -1
	n := len(s)
	res := 0
	for i := 0; i < n; i++ {
		switch s[i] {
		case 'a':
			lastIndexA = i
		case 'b':
			lastIndexB = i
		case 'c':
			lastIndexC = i
		}
		if lastIndexA != -1 && lastIndexB != -1 && lastIndexC != -1 {
			leftIndex, rightIndex := min(lastIndexA, lastIndexB, lastIndexC), max(lastIndexA, lastIndexB, lastIndexC)
			res += (leftIndex - prev) * (n - rightIndex)
			prev = leftIndex
			switch s[leftIndex] {
			case 'a':
				lastIndexA = -1
			case 'b':
				lastIndexB = -1
			case 'c':
				lastIndexC = -1
			}
		}
	}
	return res
}

func numberOfSubstrings0(s string) int {
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
