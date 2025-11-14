package p3228

func maxOperations(s string) int {
	ones := 0
	zeros := 0
	result := 0
	for _, ch := range s {
		if ch == '1' {
			if zeros > 0 {
				result += ones
			}
			ones++
			zeros = 0
		} else {
			zeros++
		}
	}
	if s[len(s)-1] == '0' {
		result += ones
	}
	return result
}

func maxOperations0(s string) int {
	ones := 0
	result := 0
	i := 0
	l := len(s)
	for i < l {
		if s[i] == '1' {
			ones++
			i++
		} else {
			result += ones
			for ; i < l && s[i] == '0'; i++ {
			}
		}
	}
	return result
}
