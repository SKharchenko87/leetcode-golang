package main

import "fmt"

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}

	res := append([]byte{'1'}, []byte(a)...)
	pref := byte(0)
	diff := len(res) - len(b)
	for j := len(b) - 1; j >= 0; j-- {
		x := res[j+diff] + b[j] + pref - 96 // 96 = '0'+'0'
		res[j+diff] = x&1 + '0'
		pref = x / 2
	}
	for i := diff - 1; i > 0 && pref > 0; i-- {
		x := res[i] - '0' + pref
		res[i] = x&1 + '0'
		pref = x / 2
	}
	return string(res[1-pref:])
}

func addBinary1(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}

	res := make([]byte, len(a)+1)
	res[0] = '1'
	pref := byte(0)
	i, j := len(a)-1, len(b)-1
	for ; i >= len(a)-len(b); i, j = i-1, j-1 {
		x := a[i] - '0' + b[j] - '0' + pref
		res[i+1] = x%2 + '0'
		pref = x / 2
	}
	for ; i >= 0; i-- {
		x := a[i] - '0' + pref
		res[i+1] = x%2 + '0'
		pref = x / 2
	}
	//res=res[1-pref:]
	return string(res[1-pref:])
}

func addBinary0(a string, b string) string {
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
