package p0409

func longestPalindrome(s string) int {
	m := make(map[int32]int, 56)
	res := 0
	odd, even := 0, 0
	for _, c := range s {
		m[c]++
		if m[c]%2 == 0 {
			even++
			odd--
			res += 2
		} else {
			if m[c] != 1 {
				even--
			}
			odd++
		}
	}
	if odd > 0 {
		res++
	}
	return res
}

func longestPalindrome0(s string) int {
	m := make(map[int32]int, 56)
	res := 0
	flgOdd := 0
	for _, c := range s {
		m[c]++
	}
	for _, v := range m {
		res += v
		if v%2 != 0 {
			res--
			flgOdd = 1
		}
	}
	return res + flgOdd
}
