package p0696

func countBinarySubstrings(s string) int {
	count := 1
	n := len(s)
	res := 0
	prevCount := 0
	for i := 1; i < n; i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			res += min(count, prevCount)
			prevCount = count
			count = 1
		}
	}
	res += min(count, prevCount)
	return res
}
