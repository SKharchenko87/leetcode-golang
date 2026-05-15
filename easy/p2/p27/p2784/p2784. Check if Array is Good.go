package p2784

import "math"

func isGood(nums []int) bool {
	l := len(nums)

	masks := [4]uint64{}
	for i := 0; i <= (l-1)/64; i++ {
		if (i+1)*64 < (l - 1) {
			masks[i] = math.MaxUint64
		} else {
			masks[i] = 1<<(l%64) - 1
		}
	}
	masks[0] >>= 1
	masks[0] <<= 1

	b := [4]uint64{}
	lastCount := 0
	for i := 0; i < l; i++ {
		high := nums[i] / 64
		low := nums[i] % 64
		b[high] |= (1 << low)
		if nums[i] == l-1 {
			lastCount++
		}
	}
	if lastCount != 2 {
		return false
	}

	for i := 0; i < 4; i++ {
		if b[i]^masks[i] != 0 {
			return false
		}
	}

	return true
}
