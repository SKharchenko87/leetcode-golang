package p0326

func isPowerOfThree(n int) bool {
	return n > 0 && 1162261467%n == 0
}

const MaxPower3 = 1162261467

func isPowerOfThree1(n int) bool {
	return n > 0 && MaxPower3%n == 0
}

var power3 = map[int]bool{1: true,
	3:          true,
	9:          true,
	27:         true,
	81:         true,
	243:        true,
	729:        true,
	2187:       true,
	6561:       true,
	19683:      true,
	59049:      true,
	531441:     true,
	177147:     true,
	1594323:    true,
	4782969:    true,
	14348907:   true,
	43046721:   true,
	129140163:  true,
	387420489:  true,
	1162261467: true,
}

func isPowerOfThree0(n int) bool {
	return power3[n]
}
