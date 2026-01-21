package p3315

import "math/bits"

func minBitwiseArray(nums []int) []int {
	mask := 1<<32 - 1
	for i := range nums {
		if nums[i] == 2 {
			nums[i] = -1
		} else {
			j := bits.TrailingZeros32(uint32(nums[i]^mask)) - 1
			nums[i] &= mask ^ (1 << j)
		}
	}
	return nums
}
