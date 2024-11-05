package p3163

func compressedString(word string) string {
	l := len(word)
	res := []byte{}
	prev := word[0]
	var cnt byte = 1
	var change = func(newCnt byte, index int) {
		res = append(res, '0'+cnt)
		res = append(res, prev)
		prev = word[index]
		cnt = newCnt
	}
	for i := 1; i < l; i++ {
		if prev == word[i] {
			if cnt == 9 {
				change(1, i)
			} else {
				cnt++
			}
		} else {
			change(1, i)
		}
	}
	change(1, 0)
	return string(res)
}

func compressedString0(word string) string {
	l := len(word)
	size := max(52, l/9*2)
	res := make([]byte, 0, size)
	prev := word[0]
	var cnt byte = 1
	var change = func(newCnt byte, index int) {
		res = append(res, '0'+cnt)
		res = append(res, prev)
		prev = word[index]
		cnt = newCnt
	}
	for i := 1; i < l; i++ {
		if prev == word[i] {
			if cnt == 9 {
				change(1, i)
			} else {
				cnt++
			}
		} else {
			change(1, i)
		}
	}
	change(1, 0)
	return string(res)
}
