package p2516

func takeCharacters(s string, k int) int {
	if k == 0 {
		return 0
	}
	l := len(s)
	var alphabet [3]int
	var leftIndex, rightIndex int
	check := func() bool {
		return alphabet[0] >= k && alphabet[1] >= k && alphabet[2] >= k
	}
	for rightIndex = l - 1; rightIndex >= 0 && !check(); rightIndex-- {
		alphabet[s[rightIndex]-'a']++
	}
	if !check() {
		return -1
	}
	ans := l - rightIndex + leftIndex
	rightIndex++
	for ; rightIndex < l && leftIndex < ans; rightIndex++ {
		alphabet[s[rightIndex]-'a']--
		for ; !check() && leftIndex < ans && leftIndex <= rightIndex; leftIndex++ {
			alphabet[s[leftIndex]-'a']++
		}
		if l-rightIndex+leftIndex < ans {
			ans = l - rightIndex + leftIndex
		}
	}

	return ans - 1
}

func takeCharacters0(s string, k int) int {
	if k == 0 {
		return 0
	}
	l := len(s)
	var alphabet [3]int
	var leftIndex, rightIndex int
	check := func() bool {
		return alphabet[0] >= k && alphabet[1] >= k && alphabet[2] >= k
	}
	for rightIndex = l - 1; rightIndex >= 0 && !check(); rightIndex-- {
		alphabet[s[rightIndex]-'a']++
	}
	if !check() {
		return -1
	}
	ans := l
	rightIndex++
	for ; rightIndex < l && leftIndex < ans; rightIndex++ {
		if l-rightIndex+leftIndex < ans {
			ans = l - rightIndex + leftIndex
		}
		alphabet[s[rightIndex]-'a']--
		for ; !check() && leftIndex < ans && leftIndex <= rightIndex; leftIndex++ {
			alphabet[s[leftIndex]-'a']++
		}
	}
	if l-rightIndex+leftIndex < ans {
		ans = l - rightIndex + leftIndex
	}

	return ans
}
