package p2825

func canMakeSubsequence(str1 string, str2 string) bool {
	l1, l2, j := len(str1), len(str2), 0
	for i := 0; i < l1 && j < l2; i++ {
		if (str2[j]+26-str1[i])%26 <= 1 {
			j++
		}
	}
	return j == l2
}

func canMakeSubsequence0(str1 string, str2 string) bool {
	l1, l2 := len(str1), len(str2)
	j := 0
	for i := 0; i < l1 && j < l2; i++ {
		if str1[i] == str2[j] || (str1[i]-'a'+27)%26 == str2[j]-'a' {
			j++
		}
	}
	return j == l2
}
