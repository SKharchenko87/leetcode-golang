package p1544

func makeGood(s string) string {
	res := []byte{}
	for i := 0; i < len(s); i++ {
		if len(res) > 0 && res[len(res)-1] == revertCase(s[i]) {
			res = res[:len(res)-1]
		} else {
			res = append(res, s[i])
		}
	}
	return string(res)
}

func revertCase(c uint8) uint8 {
	if c <= 'Z' {
		return c + 32
	}
	return c - 32
}
