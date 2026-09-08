package p3870

func countCommas(n int) int {
	if n < 1000 {
		return 0
	}
	return n%1000 + (n/1000-1)*1000 + 1
}
