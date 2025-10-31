package p3461

func hasSameDigits(s string) bool {
	l := len(s)
	tmp := []uint8(s)
	for i := 0; i < l-2; i++ {
		for j := 0; j < l-i-1; j++ {
			tmp[j] = (tmp[j] + tmp[j+1]) % 10
		}
	}
	return tmp[0] == tmp[1]
}

func hasSameDigits0(s string) bool {
	l := len(s)
	tmp := make([]uint8, l)
	for i := 0; i < l; i++ {
		tmp[i] = s[i] - '0'
	}

	for i := 0; i < l-2; i++ {
		for j := 0; j < l-i-1; j++ {
			tmp[j] = (tmp[j] + tmp[j+1]) % 10
		}
	}
	return tmp[0] == tmp[1]
}
