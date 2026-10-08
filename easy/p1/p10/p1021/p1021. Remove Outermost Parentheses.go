package p1021

func removeOuterParentheses(s string) string {
	n := len(s)
	res := make([]byte, 0, n)
	cnt := 0
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			if cnt > 0 {
				res = append(res, '(')
			}
			cnt++
		} else {
			cnt--
			if cnt > 0 {
				res = append(res, ')')
			}
		}
	}
	return string(res)
}
