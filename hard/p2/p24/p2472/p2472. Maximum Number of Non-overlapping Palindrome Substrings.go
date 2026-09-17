package p2472

func maxPalindromes(s string, k int) int {
	res := 0
	for start := 0; start < len(s)-k+1; start++ {
		if isPalindrome(s, start, start+k-1) {
			res++
			start = start + k - 1
		} else if isPalindrome(s, start, start+k) {
			res++
			start = start + k
		}
	}
	return res
}

func isPalindrome(s string, start, end int) bool {
	if len(s) <= end {
		return false
	}
	for i := 0; i <= (end-start)/2; i++ {
		if s[start+i] != s[end-i] {
			return false
		}
	}
	return true
}
