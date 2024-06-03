package p1208

func equalSubstring(s string, t string, maxCost int) int {
	startIndex, endIndex := 0, 0
	l := len(s)
	arr := []byte(s)
	tmp := 0
	res := -1
	for ; endIndex < l; endIndex++ {
		arr[endIndex] = abs(arr[endIndex], t[endIndex])
		tmp += int(arr[endIndex])
		for ; tmp > maxCost && startIndex < endIndex; startIndex++ {
			tmp -= int(arr[startIndex])
		}
		if res < endIndex-startIndex && tmp <= maxCost {
			res = endIndex - startIndex
		}
	}
	return res + 1
}

func abs(a, b byte) byte {
	if a < b {
		return b - a
	} else {
		return a - b
	}
}

func equalSubstring0(s string, t string, maxCost int) int {
	startIndex, endIndex := 0, 0
	l := len(s)
	tmp := 0
	res := -1
	for ; endIndex < l; endIndex++ {
		tmp += abs0(s[endIndex], t[endIndex])
		for ; tmp > maxCost && startIndex < endIndex; startIndex++ {
			tmp -= abs0(s[startIndex], t[startIndex])
		}
		if res < endIndex-startIndex && tmp <= maxCost {
			res = endIndex - startIndex
		}
	}
	return res + 1
}

func abs0(a, b byte) int {
	if a < b {
		return int(b - a)
	} else {
		return int(a - b)
	}
}
