package p0679

import "math"

func judgePoint24(cards []int) bool {
	var a, b, c, d float64

	a = float64(cards[0])
	b = float64(cards[1])
	c = float64(cards[2])
	d = float64(cards[3])
	if run2(a, b, c, d) {
		return true
	}
	if run2(a, b, d, c) {
		return true
	}
	if run2(a, c, b, d) {
		return true
	}
	if run2(a, c, d, b) {
		return true
	}
	if run2(a, d, b, c) {
		return true
	}
	if run2(a, d, c, b) {
		return true
	}

	if run2(b, a, c, d) {
		return true
	}
	if run2(b, a, d, c) {
		return true
	}
	if run2(b, c, a, d) {
		return true
	}
	if run2(b, c, d, a) {
		return true
	}
	if run2(b, d, a, c) {
		return true
	}
	if run2(b, d, c, a) {
		return true
	}

	if run2(c, a, b, d) {
		return true
	}
	if run2(c, a, d, b) {
		return true
	}
	if run2(c, b, a, d) {
		return true
	}
	if run2(c, b, d, a) {
		return true
	}
	if run2(c, d, a, b) {
		return true
	}
	if run2(c, d, b, a) {
		return true
	}

	if run2(d, a, b, c) {
		return true
	}
	if run2(d, a, c, b) {
		return true
	}
	if run2(d, b, a, c) {
		return true
	}
	if run2(d, b, c, a) {
		return true
	}
	if run2(d, c, a, b) {
		return true
	}
	if run2(d, c, b, a) {
		return true
	}

	return false
}

func run2(a, b, c, d float64) bool {
	var operators = [4]byte{'+', '-', '/', '*'}
	for _, oAB := range operators {
		for _, oBC := range operators {
			for _, oCD := range operators {
				if run(a, b, c, d, oAB, oBC, oCD) {
					return true
				}
			}
		}
	}
	return false
}

func e(digitLeft float64, operation byte, digitRight float64) float64 {
	switch operation {
	case '-':
		return digitLeft - digitRight
	case '*':
		return digitLeft * digitRight
	case '/':
		return digitLeft / digitRight
	default:
		return digitLeft + digitRight
	}
}

func run(a, b, c, d float64, oAB, oBC, oCD byte) bool {
	//a b c d
	if math.Round(e(e(e(a, oAB, b), oBC, c), oCD, d)*10) == 240.0 {
		return true
	}

	//(a b) c d
	if math.Round(e(e(e(a, oAB, b), oBC, c), oCD, d)*10) == 240.0 {
		return true
	}

	//a (b c) d
	if math.Round(e(e(a, oAB, e(b, oBC, c)), oCD, d)*10) == 240.0 {
		return true
	}

	//a b (c d)
	if math.Round(e(e(a, oAB, b), oBC, e(c, oCD, d))*10) == 240.0 {
		return true
	}

	//(a b) (c d)
	if math.Round(e(e(a, oAB, b), oBC, e(c, oCD, d))*10) == 240.0 {
		return true
	}

	//(a b c) d
	if math.Round(e(e(e(a, oAB, b), oBC, c), oCD, d)*10) == 240.0 {
		return true
	}

	//((a b) c) d
	if math.Round(e(e(e(a, oAB, b), oBC, c), oCD, d)*10) == 240.0 {
		return true
	}

	//(a (b c)) d
	if math.Round(e(e(a, oAB, e(b, oBC, c)), oCD, d)*10) == 240.0 {
		return true
	}

	//a (b c d)
	if math.Round(e(a, oAB, e(e(b, oBC, c), oCD, d))*10) == 240.0 {
		return true
	}

	//a ((b c) d)
	if math.Round(e(a, oAB, e(e(b, oBC, c), oCD, d))*10) == 240.0 {
		return true
	}

	//a (b (c d))
	if math.Round(e(a, oAB, e(b, oBC, e(c, oCD, d)))*10) == 240.0 {
		return true
	}

	return false
}

/*


a b c d
4 знака +-/*
3 позиции для знакаов, т.е. 4^3=64
скобки
a b c d
(a b) c d
a (b c) d
a b (c d)
(a b) (c d)
(a b c) d
((a b) c) d
(a (b c)) d
a (b c d)
a ((b c) d)
a (b (c d))
64*11 = 704


*/

// 5 1 9 1

// 3, 7, 4, 6
