package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	str := strings.ToLower(s)
	leftIndex := 0
	rightIndex := len(s) - 1
	for leftIndex < rightIndex {
		for !unicode.IsLetter(rune(str[rightIndex])) && !unicode.IsNumber(rune(str[rightIndex])) {
			rightIndex--
			if rightIndex == -1 {
				if leftIndex == 0 {
					return true
				} else {
					return false
				}
			}
		}
		for !unicode.IsLetter(rune(str[leftIndex])) && !unicode.IsNumber(rune(str[leftIndex])) {
			leftIndex++
		}
		if str[rightIndex] != str[leftIndex] {
			return false
		}
		leftIndex++
		rightIndex--
	}
	return true
}

func main() {
	fmt.Println(unicode.IsLetter(rune("0"[0])))
	fmt.Println(isPalindrome(".,"))
	fmt.Println(isPalindrome("0P"))
}
