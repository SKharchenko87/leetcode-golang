package p3718

import (
	"math"
	"math/bits"
)

func missingMultiple(nums []int, k int) int {
	counter := [2]uint64{math.MaxUint64, math.MaxUint64}
	for i := 0; i < len(nums); i++ {
		if nums[i]%k == 0 {
			v := nums[i]/k - 1
			counter[v/64] &= ^(1 << (v % 64))
		}
	}
	if v := bits.TrailingZeros64(counter[0]); v >= 0 && v < 64 {
		return (v + 1) * k
	}
	v := bits.TrailingZeros64(counter[1])
	return (v + 1 + 64) * k
}

func missingMultiple0(nums []int, k int) int {
	counter := [101]bool{}
	for i := 0; i < len(nums); i++ {
		if nums[i]%k == 0 {
			counter[nums[i]/k] = true
		}
	}
	for i := 1; i < 101; i++ {
		if !counter[i] {
			return k * i
		}
	}
	return 101
}
