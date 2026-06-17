package p3614

func processStr(s string, k int64) byte {
	var size int64
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '*':
			if size > 0 {
				size--
			}
		case '#':
			size *= 2
		case '%':
		default:
			size++
		}
	}
	if k >= size {
		return '.'
	}
	for i := len(s) - 1; i >= 0; i-- {
		switch s[i] {
		case '*':
			size++
		case '#':
			if k+1 > (size+1)/2 {
				k -= size / 2
			}
			size = (size + 1) / 2
		case '%':
			k = size - k - 1
		default:
			if k+1 == size {
				return s[i]
			}
			size--
		}
	}
	return '.'
}
