package p3090

const Checker = 0b11

func maximumLengthSubstring(s string) int {
	var (
		m      int64
		offset byte
		j, res int
	)
	for i := 0; i < len(s); i++ {
		offset = (s[i] - 'a') * 2
		m += 1 << offset
		for ; (m>>offset)&Checker == Checker; j++ {
			m -= 1 << ((s[j] - 'a') * 2)
		}
		res = max(res, i-j+1)
	}
	return res
}
