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
