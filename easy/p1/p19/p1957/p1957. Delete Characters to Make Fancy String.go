package p1957

import "strings"

func makeFancyString(s string) string {
	l := len(s)
	res := make([]byte, 0, l)
	cnt := 1
	res = append(res, s[0])
	for i := 1; i < l; i++ {
		if s[i-1] == s[i] {
			cnt++
		} else {
			cnt = 1
		}
		if cnt < 3 {
			res = append(res, s[i])
		}
	}
	return string(res)
}

func makeFancyString3(s string) string {
	sb := strings.Builder{}
	l := len(s)
	cnt := 1
	sb.WriteByte(s[0])
	for i := 1; i < l; i++ {
		if s[i-1] == s[i] {
			cnt++
		} else {
			cnt = 1
		}
		if cnt < 3 {
			sb.WriteByte(s[i])
		}
	}
	return sb.String()
}

func makeFancyString2(s string) string {
	sb := strings.Builder{}
	var prev byte
	var cnt int
	for i := 0; i < len(s); i++ {
		if prev != s[i] {
			prev = s[i]
			cnt = 0
		}
		cnt++
		if cnt < 3 {
			sb.WriteByte(s[i])
		}
	}
	return sb.String()
}

func makeFancyString1(s string) string {
	l := len(s)
	res := make([]byte, 0, l)
	ch := byte(' ')
	cnt := 0
	for i := 0; i < l; i++ {
		if ch != s[i] {
			ch = s[i]
			cnt = 0
		}
		cnt++
		if cnt < 3 {
			res = append(res, ch)
		}
	}
	return string(res)
}

func makeFancyString0(s string) string {
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
