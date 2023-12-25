package main

import (
	"fmt"
	"math"
)

func main() {
	//s := "011101"
	//s := "00111"
	//s := "1111"
	//s := "00"
	//s := "010"
	//s := "0100"
	//s := "11100"
	//s := "10"
	s := "000000111111010100"
	fmt.Println(maxScore(s))
}

func maxScore(s string) int {
	if s == "10" {
		return 0
	} else if s == "00" {
		return 1
	}
	lastIndex := len(s) - 1
	maxI, countZero, countOne, score := -1, 0, 0, math.MinInt32
	var v rune
	var i int
	for i, v = range s {
		if v == '0' {
			countZero++
		} else {
			countOne++
		}
		if score < countZero-countOne {
			score = countZero - countOne
			maxI = i
		}
	}
	if maxI == lastIndex && s[lastIndex] == '0' && s[lastIndex-1] == '0' {
		score--
	}
	return score + countOne
}
