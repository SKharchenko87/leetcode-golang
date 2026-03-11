package p1009

import (
	"math/bits"
)

const mask = 1<<32 - 1

func bitwiseComplement(num int) int {
	if num == 0 {
		return 1
	}
	offset := bits.LeadingZeros32(uint32(num))
	return num ^ (mask >> offset)
}
