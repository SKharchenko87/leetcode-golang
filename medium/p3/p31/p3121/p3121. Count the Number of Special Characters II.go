package p3121

import "math/bits"

const Mask = 0b1111_1111_1111_1111_1111_1111_11

func numberOfSpecialChars(word string) int {
	var lower, upper uint = 0, 0
	for _, ch := range word {
		if ch <= 'Z' {
			upper |= 1 << (ch - 'A')
		} else {
			if x := uint(1 << (ch - 'a')); upper&x == 0 {
				lower |= x
			} else {
				lower &= x ^ Mask
			}
		}
	}
	return bits.OnesCount(lower & upper)
}
