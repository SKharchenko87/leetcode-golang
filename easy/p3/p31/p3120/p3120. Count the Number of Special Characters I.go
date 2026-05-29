package p3120

import (
	"math/bits"
)

func numberOfSpecialChars(word string) int {
	lower, upper := [26]bool{}, [26]bool{}
	for _, ch := range word {
		if ch <= 'Z' {
			upper[ch-'A'] = true
		} else {
			lower[ch-'a'] = true
		}
	}
	ans := 0
	for i := 0; i < 26; i++ {
		if lower[i] && upper[i] {
			ans++
		}
	}
	return ans
}

func numberOfSpecialChars0(word string) int {
	lower, upper := uint32(0), uint32(0)
	for _, ch := range word {
		if ch <= 'Z' {
			upper |= (1 << (ch - 'A'))
		} else {
			lower |= (1 << (ch - 'a'))
		}
	}
	return bits.OnesCount32(lower & upper)
}
