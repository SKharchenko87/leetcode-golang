package p2381

func shiftingLetters(s string, shifts [][]int) string {
	l := len(s)
	arr := make([]int, l+1)
	for _, shift := range shifts {
		k := 1
		if shift[2] == 0 {
			k = -1
		}
		arr[shift[0]] += k
		arr[shift[1]+1] -= k
	}
	cnt := 0
	res := []byte(s)
	for i := 0; i < l; i++ {
		cnt = (cnt+arr[i])%26 + 26
		res[i] = byte(int(res[i]-'a')+cnt)%26 + 'a'
	}
	return string(res)
}

func shiftingLetters1(s string, shifts [][]int) string {
	l := len(s)
	arr := make([]int, l)
	for _, shift := range shifts {
		k := 1
		if shift[2] == 0 {
			k = -1
		}
		arr[shift[0]] += k
		if shift[1]+1 < l {
			arr[shift[1]+1] -= k
		}
	}
	cnt := 0
	res := []byte(s)
	for i := 0; i < l; i++ {
		cnt = (cnt + arr[i]) % 26
		res[i] = byte(26+int(res[i]-'a')+cnt)%26 + 'a'
	}
	return string(res)
}

// TLE
func shiftingLetters0(s string, shifts [][]int) string {
	res := []byte(s)
	for _, shift := range shifts {
		for i := shift[0]; i <= shift[1]; i++ {
			if shift[2] == 0 {
				if res[i] == 'a' {
					res[i] = 'z'
				} else {
					res[i]--
				}
			} else {
				if res[i] == 'z' {
					res[i] = 'a'
				} else {
					res[i]++
				}
			}
		}
	}
	return string(res)
}
