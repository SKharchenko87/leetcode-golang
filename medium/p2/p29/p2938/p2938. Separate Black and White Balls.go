package p2938

func minimumSteps(s string) int64 {
	l := len(s)
	var res, zero int64
	for i := l - 1; i >= 0; i-- {
		if s[i] == '0' {
			zero++
		} else {
			res += zero
		}
	}
	return res
}
