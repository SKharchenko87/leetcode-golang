package p2124

func checkString(s string) bool {
	i := 0
	for ; i < len(s) && s[i] != 'b'; i++ {
	}
	for ; i < len(s) && s[i] != 'a'; i++ {
	}
	return i == len(s)
}
