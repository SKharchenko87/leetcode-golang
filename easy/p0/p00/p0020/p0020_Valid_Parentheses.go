package p0020

var m = map[rune]rune{'(': ')', '{': '}', '[': ']'}

func isValid(s string) bool {
	stack := make([]rune, 0, len(s)/2)
	for _, ch := range s {
		switch ch {
		case ')', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != ch {
				return false
			}
			stack = stack[:len(stack)-1]
		default:
			stack = append(stack, m[ch])
		}
	}
	return len(stack) == 0
}

func check(ss *string, c string) (res bool) {
	l := len(*ss)
	if l == 0 || string((*ss)[l-1:l]) != c {
		return false
	}
	*ss = (*ss)[0 : l-1]
	return true
}

func isValid0(s string) bool {
	ss := ""
	for i := 0; i < len(s); i++ {
		v := string(s[i])
		switch v {
		case "(", "[", "{":
			ss += v
		case ")":
			if !check(&ss, "(") {
				return false
			}
		case "]":
			if !check(&ss, "[") {
				return false
			}
		case "}":
			if !check(&ss, "{") {
				return false
			}
		}

	}
	return ss == ""
}
