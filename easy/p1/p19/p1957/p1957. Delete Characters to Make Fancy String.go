package p1957

func makeFancyString(s string) string {
	res := make([]byte, 0, len(s))
	res = append(res, s[0])
	cnt := 1
	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i] {
			cnt++
		} else {
			cnt = 1
		}
		if cnt <= 2 {
			res = append(res, s[i])
		}
	}
	return string(res)
}
