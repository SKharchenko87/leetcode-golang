package main

import (
	"fmt"
)

func isPowerOfTwo(n int) bool {
	s := 0
	if n < 0 {
		return false
	}
	for n != 0 {
		if n%2 == 1 {
			s++
			if s > 1 {
				return false
			}
		}
		n = n >> 1
	}
	return s == 1
}

func main() {
	n := -16
	fmt.Println(isPowerOfTwo(n))
}
