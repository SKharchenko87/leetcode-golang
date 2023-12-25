package main

import "fmt"

func addBinary(a string, b string) string {
	lenmax := max(len(a), len(b))
	for i := len(a); i <= lenmax; i++ {
		a = "0" + a
	}
	for i := len(b); i <= lenmax; i++ {
		b = "0" + b
	}
	s := ""
	buf := false
	j := 0
	for j < lenmax {
		if buf {
			if rune(a[len(a)-1-j]) == '1' && rune(b[len(b)-1-j]) == '1' {
				s = "1" + s
				buf = true
			} else if rune(a[len(a)-1-j]) == '0' && rune(b[len(b)-1-j]) == '1' {
				s = "0" + s
				buf = true
			} else if rune(a[len(a)-1-j]) == '1' && rune(b[len(b)-1-j]) == '0' {
				s = "0" + s
				buf = true
			} else if rune(a[len(a)-1-j]) == '0' && rune(b[len(b)-1-j]) == '0' {
				s = "1" + s
				buf = false
			}
		} else {
			if rune(a[len(a)-1-j]) == '1' && rune(b[len(b)-1-j]) == '1' {
				s = "0" + s
				buf = true
			} else if rune(a[len(a)-1-j]) == '0' && rune(b[len(b)-1-j]) == '1' {
				s = "1" + s
				buf = false
			} else if rune(a[len(a)-1-j]) == '1' && rune(b[len(b)-1-j]) == '0' {
				s = "1" + s
				buf = false
			} else if rune(a[len(a)-1-j]) == '0' && rune(b[len(b)-1-j]) == '0' {
				s = "0" + s
				buf = false
			}
		}
		j++
	}
	if buf {
		s = "1" + s
	}
	return s
}

func main() {
	a := "1010"
	b := "1011"
	//Output: "10101"
	fmt.Println(addBinary(a, b))

}
