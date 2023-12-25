package main

import (
	"fmt"
)

func isNumber(s string) bool {
	e := false
	dot := false
	preE := []int32{}
	postE := []int32{}
	m := map[int32]bool{
		'-': true,
		'+': true,
		'e': true,
		'E': true,
		'.': true,
		'0': true,
		'1': true,
		'2': true,
		'3': true,
		'4': true,
		'5': true,
		'6': true,
		'7': true,
		'8': true,
		'9': true,
	}
	if s == "+." || s == ".+" || s == "-." || s == ".-" || s == "." || s == "e" || s == "E" || s == ".E" || s == ".e" || s == "e." || s == "E." || s == "+" || s == "-" {
		return false
	}
	for i, v := range s {
		//fmt.Println(i)

		if _, ok := m[v]; !ok {
			return false
		} else {
			if e {
				if len(postE) > 0 && (v == '-' || v == '+') {
					return false
				}
				if v == '.' {
					return false
				}
				postE = append(postE, v)
			} else {
				if len(preE) > 0 && (v == '-' || v == '+') {
					return false
				}
				if v == '.' {
					if dot {
						return false
					}
					dot = true
					if len(s) == 1 || (len(s) > i+1 && (s[i+1] == '.' || (s[i+1] == 'e' && len(s) == i+2))) {
						return false
					}
					//if i > 0 && (s[i-1] == '+' || s[i-1] == '-' || s[i-1] == '.') {
					//	return false
					//}
				}
				if (v == 'E' || v == 'e') && i == 0 {
					return false
				}
				if v != 'e' {
					preE = append(preE, v)
				}
			}

			if v == 'e' || v == 'E' {
				if i == 0 {
					return false
				}
				if e {
					return false
				}
				e = true
				if i == len(s)-1 {
					return false
				}
				if i > 0 && (s[i-1] == '-' || s[i-1] == '+') {
					return false
				}
				if i == 1 && s[i-1] == '.' {
					return false
				}
				if i > 1 && s[i-1] == '.' && (s[i-2] == '+' || s[i-2] == '-') {
					return false
				}
			}
		}

	}
	if e && len(postE) == 0 {
		return false
	}
	if len(postE) == 1 && (postE[0] == '-' || postE[0] == '+') {
		return false
	}
	return true
}

func main() {
	fmt.Println(isNumber("2"))
	fmt.Println(isNumber("0089"))
	fmt.Println(isNumber("-0.1"))
	fmt.Println(isNumber("+3.14"))
	fmt.Println(isNumber("4"))
	fmt.Println(isNumber("-.9"))
	fmt.Println(isNumber("2e10"))
	fmt.Println(isNumber("-90E3"))
	fmt.Println(isNumber("3e+7"))
	fmt.Println(isNumber("+6e-1"))
	fmt.Println(isNumber("53.5e93"))
	fmt.Println(isNumber("-123.456e789"))
	fmt.Println(isNumber(".1"))
	fmt.Println(isNumber("3."))
	fmt.Println(isNumber("46.e3"))
	fmt.Println(isNumber("+.8"))
	fmt.Println("=========================")
	fmt.Println(isNumber("+.E3"))
	fmt.Println(isNumber("1a"))
	fmt.Println(isNumber("1ee"))
	fmt.Println(isNumber("1e"))
	fmt.Println(isNumber("e3"))
	fmt.Println(isNumber(".e3"))
	fmt.Println(isNumber("99e2.5"))
	fmt.Println(isNumber("--6"))
	fmt.Println(isNumber("-+3"))
	fmt.Println(isNumber("95a54e53"))
	fmt.Println(isNumber("."))
	fmt.Println(isNumber(".."))
	fmt.Println(isNumber(".1."))
	fmt.Println(isNumber("4e+"))
	fmt.Println(isNumber("+E3"))
}
