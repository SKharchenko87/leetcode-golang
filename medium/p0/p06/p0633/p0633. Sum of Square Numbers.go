package p0633

import "math"

// a^2+b^2=c^2
// a|b < c
// a<c/2 c/2< b <c

func judgeSquareSum(c int) bool {
	for a := 0; a*a <= c/2; a++ {
		b := int(math.Sqrt(float64(c - a*a)))
		if a*a+b*b == c {
			return true
		}
	}
	return false
}

func judgeSquareSum1(c int) bool {
	if c%4 == 3 {
		return false
	}
	if (c-1)%4 == 0 && c%4 == 1 {
		return true
	}
	if c <= 1 || c == int(math.Pow(math.Floor(math.Sqrt(float64(c))), 2.0)) {
		return true
	}
	cf := math.Sqrt(float64(c) / 2)
	for a := 1; a <= int(math.Ceil(cf)); a++ {
		b := int(math.Sqrt(float64(c - a*a)))
		if c-a*a == b*b {
			return true
		}
	}
	return false
}

func judgeSquareSum0(c int) bool {
	if c <= 1 || c == int(math.Pow(math.Floor(math.Sqrt(float64(c))), 2.0)) {
		return true
	}
	cf := math.Sqrt(float64(c) / 2)
	for a := 1; a <= int(math.Ceil(cf)); a++ {
		b := int(math.Sqrt(float64(c - a*a)))
		if c-a*a == b*b {
			return true
		}
	}
	return false
}
