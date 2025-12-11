package p1925

import "math"

func countTriples(n int) int {
	res := 0
	for a := 1; a <= n; a++ {
		for b := a + 1; b <= n; b++ {
			c := int(math.Sqrt(float64(a*a + b*b)))
			if c <= n && a*a+b*b == c*c {
				res++
			}
		}
	}
	return res * 2
}

func countTriples0(n int) int {
	res := 0
	for a := 1; a <= n; a++ {
		for b := a + 1; b <= n; b++ {
			for c := b + 1; c <= n; c++ {
				if a*a+b*b == c*c {
					res++
				}
			}
		}
	}
	return res * 2
}
