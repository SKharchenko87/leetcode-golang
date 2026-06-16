package p3612

import "slices"

func processStr(s string) string {
	size := 0
	for i := range s {
		switch s[i] {
		case '*':
			if size > 0 {
				size--
			}
		case '#':
			size *= 2
		case '%':

		default:
			size++
		}
	}
	res := make([]byte, 0, size)
	for _, ch := range s {
		switch ch {
		case '*':
			if l := len(res); l > 0 {
				res = res[:l-1]
			}
		case '#':
			res = append(res, res...)
		case '%':
			slices.Reverse(res)
		default:
			res = append(res, byte(ch))
		}
	}
	return string(res)
}

func processStr0(s string) string {
	res := make([]byte, 0, 20)
	for _, ch := range s {
		switch ch {
		case '*':
			if l := len(res); l > 0 {
				res = res[:l-1]
			}
		case '#':
			res = append(res, res...)
		case '%':
			slices.Reverse(res)
		default:
			res = append(res, byte(ch))
		}
	}
	return string(res)
}
