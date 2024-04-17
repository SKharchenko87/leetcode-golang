package p0402

func removeKdigits(num string, k int) string {
	l := len(num)
	if k >= l {
		return "0"
	}
	res := make([]byte, l)
	resIndex := 0
	res[resIndex] = num[0]
	for i := 1; i < l; i++ {
		cur := num[i]
		for resIndex >= 0 && k > 0 && res[resIndex] > cur {
			resIndex--
			k--
		}
		if !(resIndex == -1 && cur == '0') {
			resIndex++
			res[resIndex] = cur
		}
	}
	resIndex = resIndex - k
	if resIndex < 0 {
		return "0"
	}
	return string(res[:resIndex+1])
}
