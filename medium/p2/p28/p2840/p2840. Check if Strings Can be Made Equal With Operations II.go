package p2840

func check(s1 string, s2 string, start int) bool {
	sum := [26]int{}
	for i := start; i < len(s1); i += 2 {
		sum[s1[i]-'a']++
		sum[s2[i]-'a']--
	}
	for i := 0; i < 26; i++ {
		if sum[i] != 0 {
			return false
		}
	}
	return true
}

func checkStrings(s1 string, s2 string) bool {
	return check(s1, s2, 0) && check(s1, s2, 1)
}

func checkStrings0(s1 string, s2 string) bool {
	n := len(s1)
	even, odd := [26]int{}, [26]int{}
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			even[s1[i]-'a']++
			even[s2[i]-'a']--
		} else {
			odd[s1[i]-'a']++
			odd[s2[i]-'a']--
		}
	}
	for i := 0; i < 26; i++ {
		if even[i] != 0 || odd[i] != 0 {
			return false
		}
	}
	return true
}
