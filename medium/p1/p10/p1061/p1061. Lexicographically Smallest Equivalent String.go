package p1061

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	arr := [26]int{}

	for i := range s1 {
		ch1 := s1[i] - 'a'
		ch2 := s2[i] - 'a'

		arr[ch1] |= 1 << ch1
		arr[ch1] |= 1 << ch2

		arr[ch2] |= 1 << ch1
		arr[ch2] |= 1 << ch2
	}

	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			if arr[i]&arr[j] > 0 {
				arr[i] |= arr[j]
			}
		}
	}

	l := len(baseStr)
	resBytes := []byte(baseStr)
	for i := 0; i < l; i++ {
		ch := baseStr[i] - 'a'
		for j := 0; j < 26; j++ {
			if arr[ch]&(1<<j) > 0 {
				resBytes[i] = byte(j) + 'a'
				break
			}
		}
	}

	return string(resBytes)
}

func smallestEquivalentString1(s1 string, s2 string, baseStr string) string {
	arr := [26]int{}

	for i := range s1 {
		ch1 := s1[i] - 'a'
		ch2 := s2[i] - 'a'

		arr[ch1] |= 1 << ch1
		arr[ch1] |= 1 << ch2

		arr[ch2] |= 1 << ch1
		arr[ch2] |= 1 << ch2
	}

	minX := map[int]byte{}
	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			if arr[i]&arr[j] > 0 {
				arr[i] |= arr[j]
			}
		}
		if _, ok := minX[arr[i]]; !ok {
			j := 0
			for ; j < 26 && arr[i]&(1<<j) == 0; j++ {
			}
			minX[arr[i]] = byte(j) + 'a'
		}

	}

	l := len(baseStr)
	resBytes := []byte(baseStr)
	for i := 0; i < l; i++ {

		ch := baseStr[i] - 'a'
		if arr[ch] != 0 {
			resBytes[i] = minX[arr[ch]]
		}

	}

	return string(resBytes)
}
