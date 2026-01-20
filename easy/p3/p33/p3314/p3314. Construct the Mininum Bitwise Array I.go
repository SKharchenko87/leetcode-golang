package p3314

import "math/bits"

func minBitwiseArray(nums []int) []int {
	mask := 0b1111111111
	for i := range nums {
		if nums[i] == 2 {
			nums[i] = -1
		} else {
			j := bits.TrailingZeros32(uint32(mask^nums[i])) - 1
			nums[i] &= mask ^ (1 << j)
		}
	}
	return nums
}

func minBitwiseArray0(nums []int) []int {
	l := len(nums)
	res := make([]int, l)
	mask := 0b1111111111
	for i := 0; i < l; i++ {
		if nums[i] == 2 {
			res[i] = -1
		} else {
			j := bits.TrailingZeros(uint(mask^nums[i])) - 1
			res[i] = nums[i] & (mask ^ (1 << j))
		}
	}
	return res
}
