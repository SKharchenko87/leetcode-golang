package p3304

var result = make([]byte, 0, 501)

func init() {
	result = append(result, 'a')
	for len(result) < 501 {
		for _, b := range result {
			if b == 'z' {
				result = append(result, 'a')
			} else {
				result = append(result, b+1)
			}
		}
	}
}

func kthCharacter(k int) byte {
	return result[k-1]
}
