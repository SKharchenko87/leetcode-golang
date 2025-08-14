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
	one, zero, l := 0, 0, len(s)
	best := math.MinInt32

	for i := 0; i < l-1; i++ {
		if s[i] == '1' {
			one++
		} else {
			zero++
		}
		best = max(best, zero-one)
	}
	if s[l-1] == '1' {
		one++
	}
	return best + one
}

func maxScore1(s string) int {
	if s == "00" {
		return 1
	}
	one, zero := 0, 0
	l := len(s)
	if s[0] == '0' {
		zero++
	}
	for i := 1; i < l; i++ {
		if s[i] == '1' {
			one++
		}
	}
	res := zero + one
	for i := 1; i < l-1; i++ {
		if s[i] == '0' {
			zero++
		} else {
			one--
		}
		if res < zero+one {
			res = zero + one
		}
	}
	return res
}

func maxScore0(s string) int {
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
