package p2864

//ToDo banchmark

func maximumOddBinaryNumber(s string) string {
	l := len(s)
	res := make([]byte, l+1)
	index0, index1 := l-1, 0
	for i := 0; i < l; i++ {
		if s[i] == '1' {
			res[index1] = '1'
			index1++
		} else {
			res[index0] = '0'
			index0--
		}
	}
	res[l] = '1'
	return string(res[1:])
}

func maximumOddBinaryNumber0(s string) string {
	l := len(s)
	res := make([]byte, l)
	index0, index1 := l-2, l-1
	for i := 0; i < l; i++ {
		if s[i] == '0' {
			res[index0] = '0'
			index0--
		} else {
			res[index1] = '1'
			index1 = (index1 + 1) % l
		}
	}
	return string(res)
}
