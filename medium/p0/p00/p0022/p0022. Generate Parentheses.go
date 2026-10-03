package p0022

import "strings"

func generateParenthesis(n int) []string {
	res := make([]string, 0, (n-1)*n)
	sb := strings.Builder{}
	for i := uint16(1<<n - 1); i < 1<<(2*n)-1; i++ {
		sb.Reset()
		cnt := 0
		x := i
		for j := 0; j < 2*n; j, x = j+1, x>>1 {
			if x&1 == 1 {
				cnt++
				sb.WriteRune('(')
			} else {
				cnt--
				if cnt < 0 {
					break
				}
				sb.WriteRune(')')
			}
		}
		if cnt != 0 {
			continue
		}
		res = append(res, sb.String())
	}
	return res
}

func generateParenthesis0(n int) []string {
	res := make([]string, 0, (n-1)*n)
	for i := uint16(1<<n - 1); i < 1<<(2*n)-1; i++ {
		if isValid(i, n) {
			res = append(res, numberToParenthesisString(i, n))
		}
	}
	return res
}

func isValid(x uint16, n int) bool {
	var cnt int8
	for i := 0; i < 2*n; i++ {
		if x&1 == 1 {
			cnt++
		} else {
			cnt--
			if cnt < 0 {
				return false
			}
		}
		x = x >> 1
	}
	if cnt != 0 {
		return false
	}
	return true
}

func numberToParenthesisString(x uint16, n int) string {
	res := strings.Builder{}
	for i := 0; i < 2*n; i++ {
		if x&1 == 1 {
			res.WriteRune('(')
		} else {
			res.WriteRune(')')
		}
		x >>= 1
	}
	return res.String()
}
