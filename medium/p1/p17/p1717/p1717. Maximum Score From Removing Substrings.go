package p1717

func maximumGain(s string, x int, y int) int {
	res := 0
	b := []byte(s)
	if x >= y {
		res += ab(&b) * x
		res += ba(&b) * y
	} else {
		res += ba(&b) * y
		res += ab(&b) * x
	}
	return res
}

func ab(b *[]byte) int {
	cnt := 0
	for i := 1; i < len(*b); i++ {
		if (*b)[i-1] == 'a' && (*b)[i] == 'b' {
			cnt++
			*b = append((*b)[:i-1], (*b)[i+1:]...)
			i -= 2
			if i <= 0 {
				i = 0
			}
		}
	}
	return cnt
}

func ba(b *[]byte) int {
	cnt := 0
	for i := 1; i < len(*b); i++ {
		if (*b)[i-1] == 'b' && (*b)[i] == 'a' {
			cnt++
			*b = append((*b)[:i-1], (*b)[i+1:]...)
			i -= 2
			if i <= 0 {
				i = 0
			}
		}
	}
	return cnt
}
