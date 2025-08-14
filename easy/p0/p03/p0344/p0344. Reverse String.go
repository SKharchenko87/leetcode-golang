package p0344

func reverseString(s []byte) {
	l := len(s)
	for i := 0; i < l/2; i++ {
		s[i], s[l-i-1] = s[l-i-1], s[i]
	}
}

func reverseWords(s []byte) []byte {
	reverseString(s)
	return s
}
