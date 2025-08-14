package p2220

import "math/bits"

func minBitFlips(start int, goal int) int {
	xor := start ^ goal
	res := 0
	for xor > 0 {
		xor = xor & (xor - 1)
		res++
	}
	return res
}

func minBitFlips0(start int, goal int) int {
	return bits.OnesCount32(uint32(start ^ goal))
}
