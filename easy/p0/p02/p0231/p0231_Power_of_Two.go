package p0231

// ToDo banchmark
func isPowerOfTwo(n int) bool {
	switch n {
	case 1, 1 << 1, 1 << 2, 1 << 3, 1 << 4, 1 << 5, 1 << 6, 1 << 7, 1 << 8, 1 << 9, 1 << 10, 1 << 11, 1 << 12, 1 << 13, 1 << 14, 1 << 15, 1 << 16, 1 << 17, 1 << 18, 1 << 19, 1 << 20, 1 << 21, 1 << 22, 1 << 23, 1 << 24, 1 << 25, 1 << 26, 1 << 27, 1 << 28, 1 << 29, 1 << 30, 1 << 31:
		return true
	}
	return false
}

func isPowerOfTwo2(n int) bool {
	if n <= 0 {
		return false
	}
	x := 1
	for n > x {
		x = x << 1
	}
	return n == x
}

func isPowerOfTwo1(n int) bool {
	s := 0
	if n < 0 {
		return false
	}
	for n != 0 {
		if n%2 == 1 {
			s++
			if s > 1 {
				return false
			}
		}
		n = n >> 1
	}
	return s == 1
}
