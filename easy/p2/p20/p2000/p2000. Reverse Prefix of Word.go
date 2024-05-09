package p2000

// ToDo benchmark
func reversePrefix(word string, ch byte) string {
	for i, c := range word {
		if byte(c) == ch {
			res := []byte(word)
			for j := 0; j <= i/2; j++ {
				res[j], res[i-j] = res[i-j], res[j]
			}
			return string(res)
		}
	}
	return word
}

func reversePrefix1(word string, ch byte) string {
	for i := 0; i < len(word); i++ {
		if word[i] == ch {
			res := []byte(word)
			for j := 0; j <= i/2; j++ {
				res[j], res[i-j] = res[i-j], res[j]
			}
			return string(res)
		}
	}
	return word
}

func reversePrefix0(word string, ch byte) string {
	i, l := 0, len(word)
	for ; i < l; i++ {
		if word[i] == ch {
			tmp := make([]byte, i+1)
			for j := 0; j <= i; j++ {
				tmp[j] = word[i-j]
			}
			return string(tmp) + word[i+1:]
		}
	}
	return word
}
