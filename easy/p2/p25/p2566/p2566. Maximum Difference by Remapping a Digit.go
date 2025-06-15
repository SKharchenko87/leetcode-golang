package p2566

import (
	"slices"
	"strconv"
	"strings"
)

var pow10 = make([]int, 0, 9)

func init() {
	x := 1
	for i := 0; i < 9; i++ {
		pow10 = append(pow10, x)
		x *= 10
	}
}

func minMaxDifference(num int) int {

	minX := -1
	maxX := -1
	index, _ := slices.BinarySearch(pow10, num+1)

	for i := 0; i < index; i++ {
		x := num / pow10[index-i-1] % 10
		if minX == -1 {
			minX = x
		}
		if maxX == -1 && x != 9 {
			maxX = x
		}
	}

	a, b := 0, 0
	for i := 0; i < index; i++ {
		x := num / pow10[index-i-1] % 10
		if x == minX {
			a = a*10 + 0
		} else {
			a = a*10 + x
		}
		if x == maxX {
			b = b*10 + 9
		} else {
			b = b*10 + x
		}
	}

	return b - a
}

func minMaxDifference1(num int) int {
	s := strconv.Itoa(num)
	t := s
	pos := 0
	for pos < len(s) && s[pos] == '9' {
		pos++
	}
	if pos < len(s) {
		s = strings.ReplaceAll(s, string(s[pos]), "9")
	}
	t = strings.ReplaceAll(t, string(t[0]), "0")
	max, _ := strconv.Atoi(s)
	min, _ := strconv.Atoi(t)
	return max - min
}
