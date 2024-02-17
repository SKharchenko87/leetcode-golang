package p2108

func firstPalindrome(words []string) string {
	for _, w := range words {
		if idPalindromic(w) {
			return w
		}
	}
	return ""
}

func idPalindromic(w string) bool {
	l := len(w)
	for i := 0; i < l/2; i++ {
		if w[i] != w[l-1-i] {
			return false
		}
	}
	return true
}
