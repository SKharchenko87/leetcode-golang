package p3622

func checkDivisibility(n int) bool {
	t := n
	p, s := 1, 0
	for n > 0 {
		d := n % 10
		p *= d
		s += d
		n /= 10
	}
	return t%(s+p) == 0
}
