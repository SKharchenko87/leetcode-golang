package p0342

var m = make(map[int]struct{}, 16)

func init() {
	x := 1
	for x < 2<<32 {
		m[x] = struct{}{}
		x *= 4
	}
}

func isPowerOfFour(n int) bool {
	_, exists := m[n]
	return exists
}

func isPowerOfFour1(n int) bool {
	if n <= 0 {
		return false
	}

	for n%4 == 0 {
		n /= 4
	}
	return n == 1
}

func isPowerOfFour0(n int) bool {
	if n <= 0 {
		return false
	}
	for i := 0; n > 1 && i < 32; i++ {

		if n%4 != 0 {
			return false
		}
		n /= 4
	}
	return true
}
