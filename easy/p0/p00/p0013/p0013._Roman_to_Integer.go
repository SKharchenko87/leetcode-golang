package main

import "fmt"

var roman map[string]int = map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
var roman_digit map[byte]int = map[byte]int{73: 1, 86: 5, 88: 10, 76: 50, 67: 100, 68: 500, 77: 1000}

func romanToInt(s string) int {
	prev := 0
	sum := 0
	for i := len(s) - 1; i >= 0; i-- {
		curr := roman_digit[s[i]]
		if prev > curr {
			sum -= curr
		} else {
			sum += curr
		}
		prev = curr
	}
	return sum
}

func main() {
	fmt.Println(romanToInt("XVII"))
}
