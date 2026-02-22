package p0762

import "math/bits"

// max 1111_0100_0010_0100_0000

var primes = map[int]struct{}{2: {}, 3: {}, 5: {}, 7: {}, 11: {}, 13: {}, 17: {}, 19: {}}

func countPrimeSetBits(left, right int) int {
	cnt := 0
	for i := uint(left); i <= uint(right); i++ {
		if _, ok := primes[bits.OnesCount(i)]; ok {
			cnt++
		}
	}
	return cnt
}
