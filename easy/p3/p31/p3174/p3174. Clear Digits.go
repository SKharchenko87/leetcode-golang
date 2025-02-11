package p3174

func clearDigits(s string) string {
	l := len(s)
	res := make([]byte, l)
	index := 0
	for i := 0; i < l; i++ {
		if '0' <= s[i] && s[i] <= '9' {
			index--
		} else {
			res[index] = s[i]
			index++
		}
	}
	return string(res[:index])
}
