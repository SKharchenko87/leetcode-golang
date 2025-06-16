package p1432

import (
	"slices"
	"strconv"
	"strings"
)

func maxDiff(num int) int {
	n := num
	l := 0
	digits := make([]int, 0, 10)
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
		l++
	}
	l--

	minDigit := digits[l]
	minDigitIndex := l
	if minDigit == 1 {
		for i := l - 1; i >= 0; i-- {
			d := digits[i]
			if d > 1 {
				minDigit = d
				minDigitIndex = i
				break
			}
		}
	}

	maxDigit := digits[l]
	if maxDigit == 9 {
		for i := l - 1; i >= 0; i-- {
			d := digits[i]
			if d < 9 {
				maxDigit = d
				break
			}
		}
	}

	a, b := 0, 0
	var diff int
	if minDigitIndex == l {
		diff = 1
	}
	for i := l; i >= 0; i-- {
		d := digits[i]
		a *= 10
		b *= 10
		if d == maxDigit {
			a += 9
		} else {
			a += d
		}
		if d == minDigit {
			b += diff
		} else {
			b += d
		}
	}

	return a - b
}

var pow10 = []int{1, 10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000}

func getDigit(num int, index int) int {
	return num / pow10[index] % 10
}

func maxDiff2(num int) int {
	l, _ := slices.BinarySearch(pow10, num+1)
	l--

	firstDigit := getDigit(num, l)

	minDigit := firstDigit
	minDigitIndex := l
	if minDigit == 1 {
		for i := l - 1; i >= 0; i-- {
			d := getDigit(num, i)
			if d > 1 {
				minDigit = d
				minDigitIndex = i
				break
			}
		}
	}

	maxDigit := firstDigit
	if maxDigit == 9 {
		for i := l - 1; i >= 0; i-- {
			d := getDigit(num, i)
			if d < 9 {
				maxDigit = d
				break
			}
		}
	}

	a, b := 0, 0
	var diff int
	if minDigitIndex == l {
		diff = 1
	}
	for i := l; i >= 0; i-- {
		d := getDigit(num, i)
		a *= 10
		b *= 10
		if d == maxDigit {
			a += 9
		} else {
			a += d
		}
		if d == minDigit {
			b += diff
		} else {
			b += d
		}
	}

	return a - b
}

func maxDiff0(num int) int {
	// Define a function "change", to replace x with y in a number.
	change := func(x, y int) string {
		numStr := strconv.Itoa(num)
		var res strings.Builder
		for _, digit := range numStr {
			if int(digit-'0') == x {
				res.WriteByte(byte('0' + y))
			} else {
				res.WriteRune(digit)
			}
		}
		return res.String()
	}

	minNum := num
	maxNum := num
	// Traverse all possible replacement combinations
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			res := change(x, y)
			// Check if there are leading zeros
			if res[0] != '0' {
				res_i, _ := strconv.Atoi(res)
				minNum = min(minNum, res_i)
				maxNum = max(maxNum, res_i)
			}
		}
	}

	return maxNum - minNum
}

func maxDiff1(num int) int {
	// Replace characters in the string.
	replace := func(s string, x, y byte) string {
		return strings.ReplaceAll(s, string(x), string(y))
	}
	minNum := strconv.Itoa(num)
	maxNum := strconv.Itoa(num)
	// Find a high position and replace it with 9.
	for _, digit := range maxNum {
		if digit != '9' {
			maxNum = replace(maxNum, byte(digit), '9')
			break
		}
	}
	// Replace the most significant bit with 1
	// Or find a high-order digit that is not equal to the highest digit and
	// replace it with 0.
	for i := 0; i < len(minNum); i++ {
		digit := minNum[i]
		if i == 0 {
			if digit != '1' {
				minNum = replace(minNum, digit, '1')
				break
			}
		} else {
			if digit != '0' && digit != minNum[0] {
				minNum = replace(minNum, digit, '0')
				break
			}
		}
	}

	max, _ := strconv.Atoi(maxNum)
	min, _ := strconv.Atoi(minNum)
	return max - min
}
