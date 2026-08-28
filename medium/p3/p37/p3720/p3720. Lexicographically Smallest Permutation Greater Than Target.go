package p3720

func lexGreaterPermutation(s string, target string) string {
	alphabetCount := [26]int{}
	for _, c := range s {
		alphabetCount[c-'a']++
	}
	resBytes := make([]byte, len(s))
	i := 0

	/*Prefix*/
	for ; i < len(target) && alphabetCount[target[i]-'a'] > 0; i++ {
		resBytes[i] = target[i]
		alphabetCount[target[i]-'a']--
	}

	l := byte(26)
	if i < len(target) {
		l = checkGreatChar(target[i], alphabetCount)
	}

	/*Rollback*/
	for l == 26 && i > 0 {
		i--
		alphabetCount[resBytes[i]-'a']++
		l = checkGreatChar(target[i], alphabetCount)
	}

	if l == 26 {
		return ""
	}

	/*Перелом строки когда она становится старше по порядку*/
	resBytes[i] = l + 'a'
	alphabetCount[l]--
	i++

	for j := byte(0); j < 26; j++ {
		for alphabetCount[j] > 0 {
			resBytes[i] = j + 'a'
			alphabetCount[j]--
			i++
		}
	}

	return string(resBytes[:i])
}

func checkGreatChar(c byte, alphabetCount [26]int) byte {
	l := c - 'a' + 1
	for ; l < 26 && alphabetCount[l] == 0; l++ {
	}
	return l
}

func lexGreaterPermutation0(s string, target string) string {
	alphabetCount := [26]int{}
	for _, c := range s {
		alphabetCount[c-'a']++
	}
	resBytes := make([]byte, len(target))
	f := false
	for index, c := range target {
		if f {
			for i := byte(0); i < 26; i++ {
				if alphabetCount[i] > 0 {
					resBytes[index] = i + 'a'
					alphabetCount[i]--
					break
				}
			}
		} else {
			for i := byte(0); i < 26; i++ {
				a := (byte(c) - 'a' + i) % 26
				if alphabetCount[a] > 0 {
					resBytes[index] = a + 'a'
					alphabetCount[a]--
					f = rune(a+'a') != c
					break
				}
			}
		}

	}
	res := string(resBytes)
	if res > target {
		return res
	}
	return ""
}
