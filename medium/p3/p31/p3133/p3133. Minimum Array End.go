package p3133

func minEnd(n int, x int) int64 {
	n--
	for i := 0; n > 0; i++ {
		if x>>i&1 == 0 {
			x += n & 1 << i
			n >>= 1
		}
	}
	return int64(x)
}
