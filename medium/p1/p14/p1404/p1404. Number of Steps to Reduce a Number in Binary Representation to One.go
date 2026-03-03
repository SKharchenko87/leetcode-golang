package p1404

import "strings"

func numSteps(s string) int {
	end := strings.LastIndex(s, "1") + 1
	tailZeros := len(s) - end
	if end == 1 {
		return tailZeros
	}
	return end*2 - strings.Count(s[:end], "1") + tailZeros + 1
}

func numSteps5(s string) int {
	trainZeros := 0
	i := len(s) - 1
	for ; i > 0 && s[i] == '0'; i-- {
		trainZeros++
	}
	if i == 0 {
		return trainZeros
	}
	i++
	ones := strings.Count(s[:i], "1")
	zeros := i - ones
	return zeros*2 + trainZeros + ones + 1
}

func numSteps4(s string) int {
	ones := strings.Count(s, "1")
	zeros := len(s) - ones
	tailZeros := 0
	for i := len(s) - 1; i > 0 && s[i] == '0'; i-- {
		tailZeros++
	}
	if tailZeros == zeros {
		if ones == 1 {
			return zeros
		}
		return zeros + ones + 1
	}
	return zeros*2 - tailZeros + ones + 1
}

func numSteps3(s string) int {
	res := 0
	i := len(s) - 1
	for ; i >= 0 && s[i] != '1'; i-- {
		res++
	}
	if i == 0 {
		return res
	}
	res++
	for ; i >= 0; i-- {
		res++
		if s[i] == '0' {
			res++
		}
	}
	return res
}

func numSteps2(s string) int {
	res := 0
	hasOne := false
	for i := len(s) - 1; i >= 0; i-- {
		res++
		if s[i] == '0' {
			if hasOne {
				res++
			}
		} else {
			hasOne = true
		}
	}
	return res
}

func numSteps1(s string) int {
	n0, n1 := 0, 0
	tmp := 0
	for i := 1; i < len(s); i++ {
		if s[i] == '1' {
			n1++
			tmp = n0
		} else {
			n0++
		}
	}
	if tmp > 0 || len(s) > 1 && s[1] == '1' {
		tmp++
		tmp++
	}
	return n0 + n1 + tmp
}

func numSteps0(s string) int {
	res := 0
	i := len(s) - 1
	for ; i >= 0 && s[i] != '1'; i-- {
		res++
	}
	if i == 0 {
		return res
	}
	res++
	for ; i >= 0; i-- {
		res++
		if s[i] == '0' {
			res++
		}
	}
	return res
}

//1000111
//1001000
//100100
//10010
//1001
//1010
//101
//110
//11
//100
//10
//1
