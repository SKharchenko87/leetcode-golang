package main

import "fmt"

func check(ss *string, c string) (res bool) {
	l := len(*ss)
	if l == 0 || string((*ss)[l-1:l]) != c {
		return false
	}
	*ss = (*ss)[0 : l-1]
	return true
}

func isValid(s string) bool {
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

func main() {
	fmt.Println(isValid("(){}{}"))
}
