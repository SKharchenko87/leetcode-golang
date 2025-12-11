package p1523

//go:noinline
func countOdds(low int, high int) int {
	return (high-low)>>1 + (low&1 | high&1)
}

//go:noinline
func countOdds3(low int, high int) int {
	return (high-low)>>1 + (low|high)&1
}

//go:noinline
func countOdds2(low int, high int) int {
	return (high-low)/2 + (low|high)&1
}

//go:noinline
func countOdds1(low int, high int) int {
	return (high-low)/2 + (low|high)%2
}

//go:noinline
func countOdds0(low int, high int) int {
	return (high-low)/2 + (low%2 | high%2)
}
