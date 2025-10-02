package p0921

func minAddToMakeValid(s string) int {
	o, c := 0, 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			o++
		} else {
			if o > 0 {
				o--
				c++
			}
		}
	}
	return o + c
}

func minAddToMakeValid0(s string) int {
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == ')' && len(stack) > 0 && stack[len(stack)-1] == '(' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack)
}
