package p0647

func countSubstrings(s string) int {
	count := 0
	for i := range s {
		count += isPalindromeSubstrings(s, i, i) + isPalindromeSubstrings(s, i, i+1)
	}
	return count
}

func isPalindromeSubstrings(s string, l, r int) int {
	res := 0
	for l >= 0 && r < len(s) && s[l] == s[r] {
		res++
		l--
		r++
	}
	return res
}
