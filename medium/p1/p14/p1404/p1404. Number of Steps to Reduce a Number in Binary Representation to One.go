package p1404

func numSteps(s string) int {
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
