package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	tmp := x
	ch := 0
	for tmp > 0 {
		ch = ch*10 + tmp%10
		tmp = tmp / 10
	}
	return ch == x
}

func main() {
	fmt.Println(isPalindrome(1231))
}
