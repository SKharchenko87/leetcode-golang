package p0693

func hasAlternatingBits(n int) bool {
	x := n>>1 ^ n
	return x&(x+1) == 0
}

func hasAlternatingBits0(n int) bool {
	p := n & 1
	n >>= 1
	for n > 0 {
		if p == n&1 {
			return false
		}
		p = n & 1
		n >>= 1
	}
	return true
}
