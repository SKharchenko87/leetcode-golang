package p1400

func canConstruct(s string, k int) bool {
	l := len(s)
	if l < k {
		return false
	}
	arr := [26]int{}
	i := 0
	for ; i < l; i++ {
		arr[s[i]-'a']++
	}
	for i = 0; i < 26; i++ {
		if arr[i]%2 == 1 {
			k--
		}
	}
	return 0 <= k
}

func canConstruct1(s string, k int) bool {
	if len(s) < k {
		return false
	}
	arr := [26]int{}
	for _, v := range s {
		arr[v-'a']++
	}
	for _, v := range arr {
		if v%2 != 0 {
			k--
		}
	}
	return 0 <= k
}

func canConstruct0(s string, k int) bool {
	if len(s) < k {
		return false
	}
	arr := [26]int{}
	for _, v := range s {
		arr[v-'a']++
	}
	cntOdd := 0
	for _, v := range arr {
		if v%2 == 1 {
			cntOdd++
		}
	}
	return cntOdd <= k
}
