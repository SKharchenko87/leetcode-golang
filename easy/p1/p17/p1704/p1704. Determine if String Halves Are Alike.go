package p1704

func halvesAreAlike(s string) bool {
	l := len(s)
	var cnt, i int
	for ; i < l/2; i++ {
		switch s[i] {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			cnt++
		}
	}
	for ; i < l; i++ {
		switch s[i] {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			cnt--
		}
	}
	return cnt == 0
}
